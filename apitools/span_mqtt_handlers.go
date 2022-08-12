package apitools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/lab5e/go-spanapi/v4"
)

// handleIncoming handles incoming messages that have been published on the MQTT topic
// for the configured collection.
func (m *mqttStream) handleIncoming(p *paho.Publish) {
	msg := spanapi.OutputDataMessage{}
	err := json.Unmarshal(p.Payload, &msg)
	if err != nil {
		log.Printf("error unmarshalling output message: %v", err)
		return
	}
	m.dataChan <- msg
}

// handleOnConnectionUp is called whenever we connect to the broker.  Takes care of subscribing to
// the topic belonging to the collectionID.
func (m *mqttStream) handleOnConnectionUp(mgr *autopaho.ConnectionManager, ca *paho.Connack) {
	topic := fmt.Sprintf("%s/#", m.config.collectionID)

	log.Printf("connected to MQTT broker [%s], subscribing to [%s]", m.config.brokerAddr, topic)

	_, err := mgr.Subscribe(
		context.Background(),
		&paho.Subscribe{
			Subscriptions: map[string]paho.SubscribeOptions{
				topic: {QoS: collectionTopicQoS},
			},
		})
	if err != nil {
		log.Printf("error subscrubing to [%s] on broker [%s]: %v", topic, m.config.brokerAddr, err)
	}
}

// handleDisconnect simply logs when we get disconnected from the broker.
func (m *mqttStream) handleDisconnect(d *paho.Disconnect) {
	if d.Properties != nil {
		log.Printf("disconnected from broker [%s]: %s\n", m.config.brokerAddr, d.Properties.ReasonString)
	} else {
		log.Printf("disconnected from broker [%s]; reason code: %d\n", m.config.brokerAddr, d.ReasonCode)
	}
}

func (m *mqttStream) handleOnConnectError(e error) {
	log.Printf("error during connect to mqtt broker [%s]: %v", m.config.brokerAddr, e)
}

func (m *mqttStream) handleClientError(e error) {
	log.Printf("client error: %v", e)
}
