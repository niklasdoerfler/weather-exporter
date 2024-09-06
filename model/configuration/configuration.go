package configuration

type Configuration struct {
	WebserverPort      int
	WebserverCer       string
	WebserverKey       string
	LogLevel           string
	WebserverHTTPS     WebserverHTTPSConfiguration
	JsonExporter       JsonExporterConfiguration
	PrometheusExporter PrometheusExporterConfiguration
	InfluxDbExporter   InfluxExporterConfiguration
	Mqtt               MqttConfiguration
}
