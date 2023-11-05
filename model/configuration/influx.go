package configuration

type InfluxExporterConfiguration struct {
	Enabled     bool
	Server      string
	Port        int
	Database    string
	Measurement string
	Token       string
	Org         string
}
