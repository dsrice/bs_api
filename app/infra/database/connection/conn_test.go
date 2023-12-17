package connection_test

import (
	"app/infra/database/connection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConnectionSuite struct {
	suite.Suite
}

func (s *ConnectionSuite) TestGetConnection() {
	conn := connection.Connection{}
	err := conn.GetConnection()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), conn.Conn)
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(ConnectionSuite))
}