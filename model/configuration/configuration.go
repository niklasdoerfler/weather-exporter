package configuration

type Configuration struct {
	WebserverPort      int
	LogLevel           string
	JsonExporter       JsonExporterConfiguration
	PrometheusExporter PrometheusExporterConfiguration
	InfluxDbExporter   InfluxExporterConfiguration
}
