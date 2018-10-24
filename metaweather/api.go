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

	"github.com/go-resty/resty"
)

var client *resty.Client

// Location represents metaweather location structure
type Location struct {
	Title        string `json:"title"`         // Name of the location
	LocationType string `json:"location_type"` // City|Region / State / Province|Country|Continent
	WOEID        int    `json:"woeid"`         // Where on The Earth ID
	Distance     int    `json:"distance"`      // Distance from given lattitude and longitude in meters
}

// init initiates resty client for metaweather website
func init() {
	client = resty.New().SetHostURL("https://www.metaweather.com/api")
}

// LocationSearch searches near locations for given coordinates.
func LocationSearch(latt float64, long float64) int {
	var ls []Location
	client.R().SetQueryParam("lattlong", fmt.Sprintf("%f,%f", latt, long)).SetResult(&ls).Get("/location/search/")

	return ls[0].WOEID // return nearest location to given coordinates
}

// LocationForecast forecasts weather of the 5 next days in the given location. Location is given by on the Earth ID.
func LocationForecast(woeid int) {
}
