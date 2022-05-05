package apitools

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab5e/go-spanapi/v4"
)

type mqttParam struct {
	apiToken       string
	collectionID   string
	clientID       string
	brokerOverride string
}

// MQTTOption is an option type for MQTT streams
type MQTTOption func(opt *mqttParam) *mqttParam

// WithAPIToken sets the API token to use when streaming data
func WithAPIToken(token string) MQTTOption {
	return func(o *mqttParam) *mqttParam {
		o.apiToken = token
		return o
	}
}

// WithClientID sets the client ID to use.
func WithClientID(id string) MQTTOption {
	return func(o *mqttParam) *mqttParam {
		o.clientID = id
		return o
	}
}

// WithCollectionID sets the collection ID to use when streaming data
func WithCollectionID(collectionID string) MQTTOption {
	return func(o *mqttParam) *mqttParam {
		o.collectionID = collectionID
		return o
	}
}

// WithBrokerOverride overrides the default MQTT broker endpoint.
func WithBrokerOverride(broker string) MQTTOption {
	return func(o *mqttParam) *mqttParam {
		o.brokerOverride = broker
		return o
	}
}

// NewMQTTStream creates a new data stream that uses the Span MQTT broker
func NewMQTTStream(options ...MQTTOption) (DataStream, error) {
	opts := &mqttParam{}
	for _, opt := range options {
		opts = opt(opts)
	}
	stream := &mqttStream{
		config:   *opts,
		dataChan: make(chan spanapi.OutputDataMessage),
	}
	if err := stream.start(); err != nil {
		return nil, err
	}

	return stream, nil
}

type mqttStream struct {
	config   mqttParam
	client   mqtt.Client
	dataChan chan spanapi.OutputDataMessage
}

func (m *mqttStream) start() error {
	broker := "tls://mqtt.lab5e.com:8883"
	if m.config.brokerOverride != "" {
		broker = m.config.brokerOverride
	}
	clientID := fmt.Sprintf("spanapi_%d", rand.Int())
	if m.config.clientID != "" {
		clientID = m.config.clientID
	}
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetResumeSubs(true).
		SetAutoReconnect(true).
		SetUsername(m.config.collectionID).
		SetPassword(m.config.apiToken).
		SetClientID(clientID).
		SetOnConnectHandler(func(c mqtt.Client) {
			topicName := fmt.Sprintf("%s/#", m.config.collectionID)
			if token := c.Subscribe(topicName, 0, m.messageHandler); token.Wait() && token.Error() != nil {
				log.Printf("Couldn't subscribe to %s: %v", topicName, token.Error())
				return
			}
		})

	m.client = mqtt.NewClient(opts)
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (m *mqttStream) messageHandler(client mqtt.Client, msg mqtt.Message) {
	msgStub := spanapi.OutputDataMessage{}
	if err := json.Unmarshal(msg.Payload(), &msgStub); err != nil {
		log.Printf("Error unmarshaling output message: %v", err)
		return
	}
	m.dataChan <- msgStub
}
func (m *mqttStream) Recv() (spanapi.OutputDataMessage, error) {
	return <-m.dataChan, nil
}

func (m *mqttStream) Close() error {
	m.client.Disconnect(60000)
	return nil
}
