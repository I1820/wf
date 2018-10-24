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

	"github.com/go-resty/resty"
)

var client *resty.Client

const key = "9efe7055c41ec3cb18fddd528ec555a2" // terial dark sky key, allows up to 1,000 calls per day

// init initiates resty client for metaweather website
func init() {
	client = resty.New().SetHostURL("https://api.darksky.net/")
}

// ForecastData represents dark sky forecast data.
type ForecastData struct {
	Time               int64
	Summary            string
	Icon               string
	SunriseTime        int64
	SunsetTime         int64
	MoonPhase          float64
	TemperatureMin     float64
	TemperatureMinTime int64
	TemperatureMax     float64
	TemperatureMaxTime int64
}

// ForecastRequest returns the current weather conditions,
// a minute-by-minute forecast for the next hour (where available),
// an hour-by-hour forecast for the next 48 hours, and a day-by-day forecast for the next week.
func ForecastRequest(latt float64, long float64) error {
	_, err := client.R().SetQueryParams(map[string]string{
		"exclude": "minutely,hourly,alerts,flags",
		"lang":    "en",
		"units":   "si",
	}).SetPathParams(map[string]string{
		"lattlong": fmt.Sprintf("%f,%f", latt, long),
		"key":      key,
	}).Get("forecast/{key}/{lattlong}")
	if err != nil {
		return err
	}
	return nil
}
