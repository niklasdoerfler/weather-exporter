package main

import (
	"bresser-weather-exporter/exporter/influxDb"
	"bresser-weather-exporter/exporter/mqtt"
	"bresser-weather-exporter/exporter/prometheus"
	"bresser-weather-exporter/model"
	"bresser-weather-exporter/model/configuration"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// Endpoint which is called by the weather station for transmitting the data (This endpoint appears to
	// be hardcoded in the weather station software and is appended to the server address configured in
	// the weather station settings.)
	weatherStationApiUrl = "/weatherstation/updateweatherstation.php"
)

var (
	BuildVersion = "dev"
	BuildTime    = "-"
)

var (
	weatherData = new(model.WeatherData)
	configPath  string
	config      configuration.Configuration
)

func init() {
	flag.StringVar(&configPath, "config_path", ".", "path to search for a config.yaml")
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)

	viper.SetDefault("webserverPort", 8080)
	viper.SetDefault("loglevel", "info")
	viper.SetDefault("jsonExporter", map[string]interface{}{"enabled": true})
	viper.SetDefault("prometheusExporter", map[string]interface{}{"enabled": true})
	viper.SetDefault("influxDbExporter", map[string]interface{}{"enabled": false})

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("Error reading config file, using default values. %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	logLevel, err := log.ParseLevel(config.LogLevel)
	if err == nil {
		log.SetLevel(logLevel)
	}

	log.Info("Config loaded.")
}

func runWebserver() {
	http.HandleFunc(weatherStationApiUrl, weatherStationEventHandler)

	if config.JsonExporter.Enabled {
		http.HandleFunc("/json", jsonExporterHandler)
	}

	if config.PrometheusExporter.Enabled {
		http.Handle("/metrics", promhttp.Handler())
	}

	log.Info(fmt.Sprintf("Starting Webserver on port %d", config.WebserverPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.WebserverPort), nil); err != nil {
		log.Fatal(err)
	}
}

func weatherStationEventHandler(w http.ResponseWriter, req *http.Request) {
	keys := req.URL.Query()
	log.Debug("Received new data record from weather station: ", keys)

	parseDataRecord(keys)
	prometheus.UpdatePromGauges(weatherData)
	influxDb.WriteDataToDb(weatherData)
	mqtt.PublishData(weatherData)

	_, err := io.WriteString(w, "success")
	if err != nil {
		log.Error("Unable to write response string.")
	}
}

func parseDataRecord(keys url.Values) {
	weatherData.TemperatureOutdoorFahrenheit, weatherData.TemperatureOutdoorCelsius = parseTemperature(keys.Get("tempf"))
	weatherData.TemperatureIndoorFahrenheit, weatherData.TemperatureIndoorCelsius = parseTemperature(keys.Get("indoortempf"))
	weatherData.HumidityOutdoor = parseInteger(keys.Get("humidity"))
	weatherData.HumidityOutdoorAbsolute = float32(calculateAbsoluteHumidity(weatherData.HumidityOutdoor, float64(weatherData.TemperatureOutdoorCelsius)))
	weatherData.HumidityIndoor = parseInteger(keys.Get("indoorhumidity"))
	weatherData.HumidityIndoorAbsolute = float32(calculateAbsoluteHumidity(weatherData.HumidityIndoor, float64(weatherData.TemperatureIndoorCelsius)))
	weatherData.Uv = parseFloat(keys.Get("uv"))
	weatherData.BarometerMercury, weatherData.BarometerHektopascal = parsePressure(keys.Get("baromin"))
	weatherData.RainDailyInch, weatherData.RainDailyMillimeter = parseRain(keys.Get("dailyrainin"))
	weatherData.RainCurrentInch, weatherData.RainCurrentMillimeter = parseRain(keys.Get("rainin"))
	weatherData.DewpointFahrenheit, weatherData.DewpointCelsius = parseTemperature(keys.Get("dewptf"))
	weatherData.SolarRadiation = parseFloat(keys.Get("solarradiation"))
	weatherData.WindDirection = parseInteger(keys.Get("winddir"))
	weatherData.WindGustMilesPerHour, weatherData.WindGustKilometerPerHour = parseSpeed(keys.Get("windgustmph"))
	weatherData.WindSpeedMilesPerHours, weatherData.WindSpeedKilometerPerHour = parseSpeed(keys.Get("windspeedmph"))
	weatherData.LastRefresh = time.Now()
}

func jsonExporterHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(weatherData)
	if err != nil {
		log.Error("Unable to generate json string.")
	}
}

func main() {
	flag.Parse()
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.Infof("Hello Weather Exporter %s! ☀ (Build: %s)", BuildVersion, BuildTime)
	loadConfig()
	influxDb.SetupInfluxDb(&config)
	mqtt.SetupMqttConnection(&config)
	runWebserver()
}
