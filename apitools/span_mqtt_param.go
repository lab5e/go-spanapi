package apitools

type mqttParam struct {
	apiToken     string
	collectionID string
	clientID     string
	brokerAddr   string
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
		o.brokerAddr = broker
		return o
	}
}
