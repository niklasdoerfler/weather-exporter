package configuration

type MqttConfiguration struct {
	Enabled       bool
	BrokerAddress string
	BrokerPort    int
	Username      string
	Password      string
	ClientId      string
	TopicPrefix   string
}
