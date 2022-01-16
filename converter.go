package main

import (
	"math"
	"strconv"
)

func convertFahrenheitToCelsius(tempFahrenheit float64) float64 {
	var tempCelsius = (tempFahrenheit - 32.0) * 5.0 / 9.0
	return tempCelsius
}

func convertMercuryToHektopascal(pressureMercury float64) float64 {
	return pressureMercury / 0.029529980164712
}
func convertInchToMm(lengthInch float64) float64 {
	return lengthInch / 0.03937007874
}

func convertMphToKmh(speedMph float64) float64 {
	return speedMph * 1.609344
}

func parseInteger(record string) int {
	integer, _ := strconv.ParseInt(record, 10, 32)
	return int(integer)
}

func parseFloat(record string) float32 {
	float, _ := strconv.ParseFloat(record, 32)
	return float32(float)
}

func parseTemperature(record string) (float32, float32) {
	temperature, _ := strconv.ParseFloat(record, 32)
	temperatureMetric := convertFahrenheitToCelsius(temperature)

	return float32(math.Round(temperature*100) / 100), float32(math.Round(temperatureMetric*100) / 100)
}

func parsePressure(record string) (float32, float32) {
	pressure, _ := strconv.ParseFloat(record, 32)
	pressureMetric := convertMercuryToHektopascal(pressure)

	return float32(math.Round(pressure*100) / 100), float32(math.Round(pressureMetric*100) / 100)
}

func parseRain(record string) (float32, float32) {
	rain, _ := strconv.ParseFloat(record, 32)
	rainMetric := convertInchToMm(rain)

	return float32(math.Round(rain*100) / 100), float32(math.Round(rainMetric*100) / 100)
}

func parseSpeed(record string) (float32, float32) {
	speed, _ := strconv.ParseFloat(record, 32)
	speedMetric := convertMphToKmh(speed)

	return float32(math.Round(speed*100) / 100), float32(math.Round(speedMetric*100) / 100)
}
