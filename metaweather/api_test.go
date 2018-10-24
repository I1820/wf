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
	woeid := LocationSearch(35.807425, 51.398583)

	assert.Equal(t, woeid, 2251945)
}
