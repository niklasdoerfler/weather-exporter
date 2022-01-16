package prometheus

import (
	"bresser-weather-exporter/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	gaugeWeatherTemperatureOutdoorCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_temperature_outdoor_celsius",
		Help: "The outside temperature in degree celsius.",
	})
	gaugeWeatherTemperatureOutdoorFahrenheit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_temperature_outdoor_fahrenheit",
		Help: "The outside temperature in degree fahrenheit.",
	})
)

func UpdatePromGauges(data *model.WeatherData) {
	gaugeWeatherTemperatureOutdoorCelsius.Add(float64(data.TemperatureOutdoorCelsius))
	gaugeWeatherTemperatureOutdoorFahrenheit.Add(float64(data.TemperatureOutdoorFahrenheit))
}
