package services_test

import (
	"context"
	datastore "jira-for-peasants/db"
	"jira-for-peasants/repositories"
	"jira-for-peasants/services"
	testUtils "jira-for-peasants/test_utils"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrganizationServiceTestSuite struct {
	suite.Suite
	pgContainer         *testUtils.PostgresContainer
	organizationService *services.OrganizationService
	ctx                 context.Context
}

func (suite *OrganizationServiceTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	scripts := []string{
		"test-users.sql",
	}

	pgContainer, err := testUtils.CreatePostgresContainer(suite.ctx, scripts)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgContainer
	db := datastore.NewDBFromConnectionString(pgContainer.ConnectionString)

	organizationRepository := repositories.NewOrganizationRepository(db)
	suite.organizationService = services.NewOrganizationService(organizationRepository)
}

func (suite *OrganizationServiceTestSuite) TestCreateOrganization() {
	_, err := suite.organizationService.CreateOrganization(suite.ctx, services.CreateOrganizationParams{
		Name:   "Test Organization",
		UserId: "123123",
	})

	suite.Nil(err)
}

func (suite *OrganizationServiceTestSuite) TestCreateOrganizationWithInvalidUser() {
	_, err := suite.organizationService.CreateOrganization(suite.ctx, services.CreateOrganizationParams{
		Name:   "Invalid User",
		UserId: "INoExist",
	})

	suite.NotNil(err)
}

func (suite *OrganizationServiceTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestOrganizationTestSuite(t *testing.T) {
	suite.Run(t, new(OrganizationServiceTestSuite))
}
