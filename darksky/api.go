/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 24-10-2018
 * |
 * | File Name:     api.go
 * +===============================================
 */

package darksky

import (
	"fmt"

	"gopkg.in/resty.v1"
)

var client *resty.Client

const key = "9efe7055c41ec3cb18fddd528ec555a2" // terial dark sky key, allows up to 1,000 calls per day

// init initiates resty client for metaweather website
func init() {
	client = resty.New().SetHostURL("https://api.darksky.net/")
}

// ForecastData represents dark sky forecast data.
type ForecastData struct {
	Time    int64  `json:"time"`
	Summary string `json:"summary"`
	Icon    string `json:"icon"`

	SunriseTime int64 `json:"sunriseTime"`
	SunsetTime  int64 `json:"sunsetTime"`

	MoonPhase float64 `json:"moonPhase"`

	Temperature         float64 `json:"temperature"`
	TemperatureLow      float64 `json:"temperatureLow"`
	TemperatureLowTime  int64   `json:"temperatureLowTime"`
	TemperatureHigh     float64 `json:"temperatureHigh"`
	TemperatureHighTime int64   `json:"temperatureHighTime"`

	Humidity    float64 `json:"humidity"`
	WindBearing float64 `json:"windBearing"`
	WindSpeed   float64 `json:"windSpeed"`
}

// ForecastResponse represents dark sky forecast response that contains
// ForecastData in its fields.
type ForecastResponse struct {
	Latitude  float64      `json:"latitude"`
	Longitude float64      `json:"longitude"`
	Timezone  string       `json:"timezone"`
	Currently ForecastData `json:"currently"`
	Daily     struct {
		Summery string         `json:"summery"`
		Icon    string         `json:"icon"`
		Data    []ForecastData `json:"data"`
	} `json:"daily"`
}

// ForecastRequest returns the current weather conditions,
// a minute-by-minute forecast for the next hour (where available),
// an hour-by-hour forecast for the next 48 hours, and a day-by-day forecast for the next week.
func ForecastRequest(latt float64, long float64) (ForecastResponse, error) {
	var w ForecastResponse

	_, err := client.R().SetResult(&w).SetQueryParams(map[string]string{
		"exclude": "minutely,hourly,alerts,flags",
		"lang":    "en",
		"units":   "si",
	}).SetPathParams(map[string]string{
		"lattlong": fmt.Sprintf("%f,%f", latt, long),
		"key":      key,
	}).Get("forecast/{key}/{lattlong}")
	if err != nil {
		return w, err
	}

	return w, nil
}
