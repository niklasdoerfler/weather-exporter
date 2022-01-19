package mqtt

import (
	"bresser-weather-exporter/model"
	"bresser-weather-exporter/model/configuration"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	mqttClient  mqtt.Client
	topicPrefix string
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Debugf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	rOps := client.OptionsReader()
	servers := rOps.Servers()
	log.Infof("Connected to mqtt broker: %s:%s", servers[0].Host, servers[0].Port())
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Warnf("Connect to mqtt broker lost: %v", err)
}

func SetupMqttConnection(config *configuration.Configuration) {
	if config.Mqtt.Enabled {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Mqtt.BrokerAddress, config.Mqtt.BrokerPort))
		opts.SetClientID(config.Mqtt.ClientId)
		opts.SetUsername(config.Mqtt.Username)
		opts.SetPassword(config.Mqtt.Password)
		opts.SetDefaultPublishHandler(messagePubHandler)
		opts.OnConnect = connectHandler
		opts.OnConnectionLost = connectLostHandler
		topicPrefix = strings.TrimSuffix(config.Mqtt.TopicPrefix, "/")
		mqttClient = mqtt.NewClient(opts)
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			log.Error("Unable to connect to mqtt broker:", token.Error())
		}
	}
}

func PublishData(data *model.WeatherData) {
	if mqttClient != nil && mqttClient.IsConnectionOpen() {
		jsonData, _ := json.Marshal(data)
		publishMessage("json", string(jsonData))

		var inInterface map[string]interface{}
		err := json.Unmarshal(jsonData, &inInterface)
		if err == nil {
			for field, value := range inInterface {
				publishMessage(field, fmt.Sprintf("%v", value))
			}
		}
	}
}

func publishMessage(topic string, message string) {
	fullTopic := topicPrefix + "/" + topic
	log.Debugf("Publish message on topic %s: '%s'", fullTopic, message)
	token := mqttClient.Publish(fullTopic, 0, false, message)
	token.Wait()
}
