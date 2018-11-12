/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2018
 * |
 * | File Name:     darksky_test.go
 * +===============================================
 */

package actions

import "github.com/I1820/wf/darksky"

func (as *ActionSuite) Test_DarkskyHandler() {
	var wr darksky.ForecastResponse

	// Create (POST /api/darksky)
	res := as.JSON("/api/darksky").Post(darkskyReq{Latitude: 35.8064342, Longitude: 51.3950481})
	as.Equalf(200, res.Code, "Error: %s", res.Body.String())
	res.Bind(&wr)
}
