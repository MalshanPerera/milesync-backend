package services_test

import (
	"jira-for-peasants/services"
	testUtils "jira-for-peasants/test_utils"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	pgContainer *testUtils.PostgresContainer
	userService *services.UserService
}
