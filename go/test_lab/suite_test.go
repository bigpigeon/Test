package test_lab

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExample() {
	suite.Equal(suite.VariableThatShouldStartAtFive, 5)
	suite.T().Log("example 1")
}

func (suite *ExampleTestSuite) TestExample2() {
	suite.Equal(suite.VariableThatShouldStartAtFive, 5)
	suite.T().Log("example 2")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
