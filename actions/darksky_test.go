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

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/I1820/wf/darksky"
)

func (suite *WFTestSuite) Test_DarkskyHandler() {
	var wr darksky.ForecastResponse

	// Create (POST /api/darksky)
	w := httptest.NewRecorder()
	data, err := json.Marshal(darkskyReq{Latitude: 35.8064342, Longitude: 51.3950481})
	suite.NoError(err)
	req, err := http.NewRequest("POST", "/api/darksky", bytes.NewReader(data))
	suite.NoError(err)
	suite.engine.ServeHTTP(w, req)

	suite.Equalf(200, w.Code, "Error: %s", w.Body.String())
	suite.NoError(json.Unmarshal(w.Body.Bytes(), &wr))
}
