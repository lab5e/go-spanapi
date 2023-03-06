package apitools

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/lab5e/go-spanapi/v4"
)

type mqttStream struct {
	config   mqttParam
	client   *autopaho.ConnectionManager
	dataChan chan spanapi.OutputDataMessage
}

// defaults, these should be reasonable
const (
	keepAliveSeconds   = 30
	connectRetryDelay  = 5 * time.Second
	connectTimeout     = 5 * time.Second
	closeTimeout       = 15 * time.Second
	collectionTopicQoS = 1
	defaultBrokerAddr  = "tls://mqtt.lab5e.com:8883"
)

// errors
var (
	ErrTimeout             = errors.New("timed out")
	ErrEmptyBrokerAddress  = errors.New("broker address is empty")
	ErrMissingAPIToken     = errors.New("api token has not been set")
	ErrMissingCollectionID = errors.New("collection ID has not been set")
)

// NewMQTTStream creates a new data stream that uses the Span MQTT broker
func NewMQTTStream(options ...MQTTOption) (DataStream, error) {
	opts := &mqttParam{
		brokerAddr: defaultBrokerAddr,
	}

	for _, opt := range options {
		opts = opt(opts)
	}

	if opts.brokerAddr == "" {
		return nil, ErrEmptyBrokerAddress
	}

	if opts.apiToken == "" {
		return nil, ErrMissingAPIToken
	}

	if opts.collectionID == "" {
		return nil, ErrMissingCollectionID
	}

	stream := &mqttStream{
		config:   *opts,
		dataChan: make(chan spanapi.OutputDataMessage),
	}

	err := stream.start()
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (m *mqttStream) Recv() (spanapi.OutputDataMessage, error) {
	return <-m.dataChan, nil
}

func (m *mqttStream) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), closeTimeout)
	defer cancel()

	return m.client.Disconnect(ctx)
}

// start the MQTT broker client.  This client will auto-reconnect if the connection
// falls down.
func (m *mqttStream) start() error {
	brokerURL, err := url.Parse(m.config.brokerAddr)
	if err != nil {
		return err
	}

	buf := make([]byte, 8)
	rand.Read(buf)
	clientID := fmt.Sprintf("spanapi_%d", binary.BigEndian.Uint64(buf))
	if m.config.clientID != "" {
		clientID = m.config.clientID
	}

	clientConfig := autopaho.ClientConfig{
		BrokerUrls:        []*url.URL{brokerURL},
		KeepAlive:         keepAliveSeconds,
		ConnectRetryDelay: connectRetryDelay,
		ConnectTimeout:    connectTimeout,
		OnConnectionUp:    m.handleOnConnectionUp,
		OnConnectError:    m.handleOnConnectError,

		ClientConfig: paho.ClientConfig{
			ClientID:           clientID,
			Router:             paho.NewSingleHandlerRouter(m.handleIncoming),
			OnServerDisconnect: m.handleDisconnect,
			OnClientError:      m.handleClientError,
		},
	}

	clientConfig.SetUsernamePassword(m.config.collectionID, []byte(m.config.apiToken))

	m.client, err = autopaho.NewConnection(context.Background(), clientConfig)
	if err != nil {
		return fmt.Errorf("error creating new connection: %v", err)
	}
	return nil
}
