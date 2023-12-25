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
	conn := connection.NewConnection()

	assert.NotNil(s.T(), conn.Conn)

	_, err := conn.Conn.Query("select * from schema_migrations")

	assert.Nil(s.T(), err)
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(ConnectionSuite))
}