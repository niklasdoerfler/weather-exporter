package configuration

type InfluxExporterConfiguration struct {
	Enabled     bool
	Server      string
	Port        int
	Username    string
	Password    string
	Database    string
	Measurement string
}
