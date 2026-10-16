package configuration

type Configuration struct {
	WebserverPort      int
	LogLevel           string
	Https              HttpsConfiguration
	JsonExporter       JsonExporterConfiguration
	PrometheusExporter PrometheusExporterConfiguration
	InfluxDbExporter   InfluxExporterConfiguration
	Mqtt               MqttConfiguration
}
