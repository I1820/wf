/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 24-10-2018
 * |
 * | File Name:     api_test.go
 * +===============================================
 */

package metaweather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationSearch(t *testing.T) {
	woeid, err := LocationSearch(35.807425, 51.398583) // 18.20 coordinate
	assert.NoError(t, err)

	assert.Equal(t, woeid, 2251945) // Tehran WOEID
}

func TestLocationForecast(t *testing.T) {
	ws, err := LocationForecast(2251945) // Tehran WOEID
	assert.NoError(t, err)

	t.Log(ws)
}
