package errormessage_test

import (
	"app/infra/errormessage"
	"app/repositories/ri"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ErrorSuite struct {
	suite.Suite
	repo ri.UserRepository
}

func (s *ErrorSuite) TestLoadSuccess() {
	infos := errormessage.SettingError()

	assert.NotEqual(s.T(), len(infos), 0)
	assert.Equal(s.T(), infos[errormessage.BadRequest].Code, 1000)
	assert.Equal(s.T(), infos[errormessage.BadRequest].Status, 400)
	assert.Equal(s.T(), infos[errormessage.BadRequest].Message, "リクエストパラメータ異常")
}

func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}