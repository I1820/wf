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

package metaweather

import (
	"fmt"
	"time"

	"gopkg.in/resty.v1"
)

var client *resty.Client

// Location represents metaweather location structure
type Location struct {
	Title        string `json:"title"`         // Name of the location
	LocationType string `json:"location_type"` // City|Region / State / Province|Country|Continent
	WOEID        int    `json:"woeid"`         // Where on The Earth ID
	Distance     int    `json:"distance"`      // Distance from given latitude and longitude in meters
}

// Weather represents metaweather consolidated weathers in each day
type Weather struct {
	ID                   int       `json:"id"`
	WeatherStateName     string    `json:"weather_state_name"`
	WeatherStateAbbr     string    `json:"weather_state_abbr"`
	WindDirectionCompass string    `json:"wind_direction_compass"`
	Created              time.Time `json:"created"`
	ApplicableDate       string    `json:"applicable_date"`
	MinTemperature       float64   `json:"min_temp"`
	MaxTemperature       float64   `json:"max_temp"`
	TheTemperature       float64   `json:"the_temp"`
	WindSpeed            float64   `json:"wind_speed"`
	WindDirection        float64   `json:"wind_direction"`
	AirPressure          float64   `json:"air_pressure"`
	Humidity             int       `json:"humidity"`
	Visibility           float64   `json:"visibility"`
	Predictability       int       `json:"predictability"`
}

// init initiates resty client for metaweather website
func init() {
	client = resty.New().SetHostURL("https://www.metaweather.com/api")
}

// LocationSearch searches near locations for given coordinates.
func LocationSearch(latt float64, long float64) (int, error) {
	var ls []Location
	_, err := client.R().SetQueryParam("lattlong", fmt.Sprintf("%f,%f", latt, long)).SetResult(&ls).Get("/location/search/")
	if err != nil {
		return 0, err
	}

	return ls[0].WOEID, nil // return nearest location to given coordinates
}

// LocationForecast forecasts weather of the 5 next days in the given location. Location is given by on the Earth ID.
func LocationForecast(woeid int) ([]Weather, error) {
	var w struct {
		ConsolidatedWeather []Weather `json:"consolidated_weather"`
		Time                time.Time `json:"time"`
		SunRise             time.Time `json:"sun_rise"`
		SunSet              time.Time `json:"sun_set"`
	}

	_, err := client.R().SetResult(&w).SetPathParams(map[string]string{
		"woeid": fmt.Sprintf("%d", woeid),
	}).Get("/location/{woeid}/")
	if err != nil {
		return nil, err
	}

	return w.ConsolidatedWeather, nil
}
