package influxDb

import (
	"bresser-weather-exporter/model"
	"bresser-weather-exporter/model/configuration"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	log "github.com/sirupsen/logrus"
)

var (
	influxClient    influxdb2.Client
	influxWriteAPI  api.WriteAPIBlocking
	measurementName string
)

func SetupInfluxDb(config *configuration.Configuration) {
	if config.InfluxDbExporter.Enabled {
		log.Info("Configuring influx db...")
		influxClient = influxdb2.NewClient(
			fmt.Sprintf("http://%s:%d", config.InfluxDbExporter.Server, config.InfluxDbExporter.Port),
			fmt.Sprintf("%s:%s", config.InfluxDbExporter.Username, config.InfluxDbExporter.Password),
		)
		influxWriteAPI = influxClient.WriteAPIBlocking("", config.InfluxDbExporter.Database)
		measurementName = config.InfluxDbExporter.Measurement
	}
}

func WriteDataToDb(data *model.WeatherData) {
	if influxClient != nil {
		point := influxdb2.NewPoint(measurementName,
			map[string]string{},
			map[string]interface{}{
				"temp_indoor":     data.TemperatureIndoorCelsius,
				"temp_indoor_f":   data.TemperatureIndoorFahrenheit,
				"temp_outdoor":    data.TemperatureOutdoorCelsius,
				"temp_outdoor_f":  data.TemperatureOutdoorFahrenheit,
				"hum_indoor":      data.HumidityIndoor,
				"hum_outdoor":     data.HumidityOutdoor,
				"uv":              data.Uv,
				"barom":           data.BarometerHektopascal,
				"barom_m":         data.BarometerMercury,
				"daily_rain":      data.RainDailyMillimeter,
				"daily_rain_i":    data.RainDailyInch,
				"rain":            data.RainCurrentMillimeter,
				"rain_i":          data.RainCurrentInch,
				"dewpoint":        data.DewpointCelsius,
				"dewpoint_f":      data.DewpointFahrenheit,
				"solar_radiation": data.SolarRadiation,
				"winddir":         data.WindDirection,
				"wind_gust":       data.WindGustKilometerPerHour,
				"wind_gust_m":     data.WindGustMilesPerHour,
				"wind_speed":      data.WindSpeedKilometerPerHour,
				"wind_speed_m":    data.WindSpeedMilesPerHours,
			},
			data.LastRefresh)

		err := influxWriteAPI.WritePoint(context.Background(), point)
		if err != nil {
			log.Errorf("Write error: %s\n", err.Error())
		}
	}
}
