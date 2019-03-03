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

package darksky

import (
	"testing"

	"github.com/I1820/wf/config"
	"github.com/stretchr/testify/assert"
)

func TestForecastRequest(t *testing.T) {
	darksky := NewDarksky(config.GetConfig().Darksky.Key)

	wr, err := darksky.ForecastRequest(35.807425, 51.398583) // 18.20 coordinate
	assert.NoError(t, err)

	t.Logf("%+v", wr)
}
