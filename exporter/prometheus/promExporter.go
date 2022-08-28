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

	gaugeWeatherTemperatureIndoorCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_temperature_indoor_celsius",
		Help: "The inside temperature in degree celsius.",
	})

	gaugeWeatherTemperatureIndoorFahrenheit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_temperature_indoor_fahrenheit",
		Help: "The inside temperature in degree fahrenheit.",
	})

	gaugeWeatherHumidityOutdoor = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_humidity_outdoor",
		Help: "The outside relative humidity in percent.",
	})

	gaugeWeatherHumidityOutdoorAbsolute = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_humidity_outdoor_absolute",
		Help: "The outside absolute humidity in gram per cubic meter.",
	})

	gaugeWeatherHumidityIndoor = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_humidity_indoor",
		Help: "The inside relative humidity in percent.",
	})

	gaugeWeatherHumidityIndoorAbsolute = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_humidity_indoor_absolute",
		Help: "The inside absolute humidity in gram per cubic meter.",
	})

	gaugeWeatherUv = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_uv",
		Help: "The outside uv index.",
	})

	gaugeWeatherBarometerHektopascal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_barometer_hektopascal",
		Help: "The barometer (air pressure) in hektropascal.",
	})

	gaugeWeatherBarometerMercury = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_barometer_mercury",
		Help: "The barometer (air pressure) in mercury.",
	})

	gaugeWeatherRainDailyMillimeter = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_rain_daily_millimeter",
		Help: "The daily sum of precipitation (rain) in millimeter.",
	})

	gaugeWeatherRainDailyInch = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_rain_daily_inch",
		Help: "The daily sum of precipitation (rain) in inch.",
	})

	gaugeWeatherRainCurrentMillimeter = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_rain_current_millimeter",
		Help: "The current amount of precipitation (rain) in millimeter.",
	})

	gaugeWeatherRainCurrentInch = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_rain_current_inch",
		Help: "The current amount of precipitation (rain) in inch.",
	})

	gaugeWeatherDewpointOutdoorCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_dewpoint_outdoor_celsius",
		Help: "The outside dew-point in degree celsius.",
	})

	gaugeWeatherDewpointOutdoorFahrenheit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_dewpoint_outdoor_fahrenheit",
		Help: "The outside dew-point in degree fahrenheit.",
	})

	gaugeWeatherSolarRadiation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_solar_radiation",
		Help: "The solar radiation in watts per square meter.",
	})

	gaugeWeatherWindDirection = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_wind_direction",
		Help: "The wind direction in degree.",
	})

	gaugeWeatherWindGustKilometerPerHour = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_wind_gust_kilometer_per_hour",
		Help: "The wind gust speed in kilometers per hour.",
	})

	gaugeWeatherWindGustMilesPerHour = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_wind_gust_miles_per_hour",
		Help: "The wind gust speed in miles per hour.",
	})

	gaugeWeatherWindSpeedKilometerPerHour = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_wind_speed_kilometer_per_hour",
		Help: "The wind speed in kilometers per hour.",
	})

	gaugeWeatherWindSpeedMilesPerHour = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "weather_wind_speed_miles_per_hour",
		Help: "The wind speed in miles per hour.",
	})
)

func UpdatePromGauges(data *model.WeatherData) {
	gaugeWeatherTemperatureOutdoorCelsius.Set(float64(data.TemperatureOutdoorCelsius))
	gaugeWeatherTemperatureOutdoorFahrenheit.Set(float64(data.TemperatureOutdoorFahrenheit))
	gaugeWeatherTemperatureIndoorCelsius.Set(float64(data.TemperatureIndoorCelsius))
	gaugeWeatherTemperatureIndoorFahrenheit.Set(float64(data.TemperatureIndoorFahrenheit))
	gaugeWeatherHumidityOutdoor.Set(float64(data.HumidityOutdoor))
	gaugeWeatherHumidityOutdoorAbsolute.Set(float64(data.HumidityOutdoorAbsolute))
	gaugeWeatherHumidityIndoor.Set(float64(data.HumidityIndoor))
	gaugeWeatherHumidityIndoorAbsolute.Set(float64(data.HumidityIndoorAbsolute))
	gaugeWeatherUv.Set(float64(data.Uv))
	gaugeWeatherBarometerHektopascal.Set(float64(data.BarometerHektopascal))
	gaugeWeatherBarometerMercury.Set(float64(data.BarometerMercury))
	gaugeWeatherRainDailyMillimeter.Set(float64(data.RainDailyMillimeter))
	gaugeWeatherRainDailyInch.Set(float64(data.RainDailyInch))
	gaugeWeatherRainCurrentMillimeter.Set(float64(data.RainCurrentMillimeter))
	gaugeWeatherRainCurrentInch.Set(float64(data.RainCurrentInch))
	gaugeWeatherDewpointOutdoorCelsius.Set(float64(data.DewpointCelsius))
	gaugeWeatherDewpointOutdoorFahrenheit.Set(float64(data.DewpointFahrenheit))
	gaugeWeatherSolarRadiation.Set(float64(data.SolarRadiation))
	gaugeWeatherWindDirection.Set(float64(data.WindDirection))
	gaugeWeatherWindGustKilometerPerHour.Set(float64(data.WindSpeedKilometerPerHour))
	gaugeWeatherWindGustMilesPerHour.Set(float64(data.WindGustMilesPerHour))
	gaugeWeatherWindSpeedKilometerPerHour.Set(float64(data.WindSpeedKilometerPerHour))
	gaugeWeatherWindSpeedMilesPerHour.Set(float64(data.WindSpeedMilesPerHours))
}
