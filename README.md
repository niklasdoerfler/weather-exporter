# Bresser Weather Exporter

This project is intended to be used as a middleware between a Wi-Fi enabled weather station, using the Weather
Underground protocol. The exporter boots up a webserver listening for requests from the weather station and stores the
data in an internal data structure. It provides the measured values via different exporters (see below). Values are
provided both in metric and imperial units, when applicable.

The project has been developed using
the [Bresser 7in1 Clear View Weather Station](https://www.bresser.de/Wetter-Zeit/BRESSER-WLAN-ClearView-Wettercenter-mit-7-in-1-Profi-Sensor.html)
.

## Principle of operation

The above-mentioned weather station is WLAN capable and provides a mode in which it opens a WLAN access point, via which
a configuration page can be opened in a web browser. Here, among other things, settings for the time zone and
coordinates of the location can be configured. Furthermore, API tokens for services like Weather Underground can be
stored. In addition, a freely definable server address can be configured for third-party services.

This project takes advantage of this fact and provides a compatible server. The weather station calls the server address
stored there with a fixed configured endpoint (`/weatherstation/updateweatherstation.php`). The corresponding measured
values are transferred via GET parameters. Such a request looks like the following:
`http://192.168.1.10:8080/weatherstation/updateweatherstation.php?ID=station-id8&PASSWORD=password&action=updateraww&realtime=1&rtfreq=5&dateutc=now&baromin=30.23&tempf=37.5&dewptf=36.6&humidity=97&windspeedmph=0.0&windgustmph=0.0&winddir=306&rainin=0.09&dailyrainin=0.45&solarradiation=0.0&UV=0.0&indoortempf=65.4&indoorhumidity=50`

Since the transmitted measured values are available only in imperial format, these are converted afterwards by this
project accordingly also into metric units. The values of both unit systems are then provided.

## Exporters

### Influx DB

The measured values received by the weather station can be written into an influx database. Influx settings are set via
config file (see config example below).

### HTTP JSON

This exporter provides a simple http endpoint (`/json`) serving the latest measurement values in json format like
follows:

```json
{
  "TemperatureIndoorCelsius": 18.56,
  "TemperatureIndoorFahrenheit": 65.4,
  "TemperatureOutdoorCelsius": 3.06,
  "TemperatureOutdoorFahrenheit": 37.5,
  "HumidityIndoor": 50,
  "HumidityOutdoor": 97,
  "Uv": 0,
  "BarometerHektopascal": 1023.71,
  "BarometerMercury": 30.23,
  "RainDailyMillimeter": 11.43,
  "RainDailyInch": 0.45,
  "RainCurrentMillimeter": 2.29,
  "RainCurrentInch": 0.09,
  "DewpointCelsius": 2.56,
  "DewpointFahrenheit": 36.6,
  "SolarRadiation": 0,
  "WindDirection": 306,
  "WindGustKilometerPerHour": 0,
  "WindGustMilesPerHour": 0,
  "WindSpeedKilometerPerHour": 0,
  "WindSpeedMilesPerHours": 0,
  "LastRefresh": "2022-01-16T17:25:06.6178559+01:00"
}
```

### Prometheus

This exporter provides prometheus compatible metrics on the metrics endpoint (`/metrics`).

## Config

The service can be configured by a `config.yaml` file placed next to the binary.

The following represents an example config:

```yaml
webserverPort: 8080

logLevel: info

jsonExporter:
  enabled: true

prometheusExporter:
  enabled: true

influxDbExporter:
  enabled: false
  server: 1.2.3.4
  port: 8086
  user: user
  password: password
  database: database
  measurement: measurement

mqtt:
  enabled: false
  brokerAddress: 1.2.3.4
  brokerPort: 1883
  username: username
  password: password
  clientId: weather-exporter
  topicPrefix: weather
```