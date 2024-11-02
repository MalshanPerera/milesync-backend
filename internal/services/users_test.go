package services_test

import (
	"jira-for-peasants/internal/services"
	testUtils "jira-for-peasants/test/test_utils"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	pgContainer *testUtils.PostgresContainer
	userService *services.UserService
}
