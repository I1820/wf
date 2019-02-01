/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-02-2019
 * |
 * | File Name:     app_test.go
 * +===============================================
 */

package actions

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

// WFTestSuite is a test suite for wf component APIs.
type WFTestSuite struct {
	suite.Suite
	engine *echo.Echo
}

// SetupSuite initiates wf test suite
func (suite *WFTestSuite) SetupSuite() {
	suite.engine = App()
}

// Let's test wf APIs!
func TestService(t *testing.T) {
	suite.Run(t, new(WFTestSuite))
}
