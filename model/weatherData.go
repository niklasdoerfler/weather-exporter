package model

import "time"

type WeatherData struct {
	TemperatureIndoorCelsius     float32
	TemperatureIndoorFahrenheit  float32
	TemperatureOutdoorCelsius    float32
	TemperatureOutdoorFahrenheit float32
	HumidityIndoor               int
	HumidityOutdoor              int
	Uv                           float32
	BarometerHektopascal         float32
	BarometerMercury             float32
	RainDailyMillimeter          float32
	RainDailyInch                float32
	RainCurrentMillimeter        float32
	RainCurrentInch              float32
	DewpointCelsius              float32
	DewpointFahrenheit           float32
	SolarRadiation               float32
	WindDirection                int
	WindGustKilometerPerHour     float32
	WindGustMilesPerHour         float32
	WindSpeedKilometerPerHour    float32
	WindSpeedMilesPerHours       float32
	LastRefresh                  time.Time
}
