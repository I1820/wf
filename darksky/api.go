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

import "github.com/go-resty/resty"

var client *resty.Client

// init initiates resty client for metaweather website
func init() {
	client = resty.New().SetHostURL("https://api.darksky.net/")
}

// ForecastRequest returns the current weather conditions,
// a minute-by-minute forecast for the next hour (where available),
// an hour-by-hour forecast for the next 48 hours, and a day-by-day forecast for the next week.
func ForecastRequest() {
}
