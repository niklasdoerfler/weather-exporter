package configuration

type HttpsConfiguration struct {
	Enabled  bool
	Port     int
	CertFile string
	KeyFile  string
}
