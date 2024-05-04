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

type ProjectServiceTestSuite struct {
	suite.Suite
	pgContainer    *testUtils.PostgresContainer
	projectService *services.ProjectService
	ctx            context.Context
}

func (suite *ProjectServiceTestSuite) SetupSuite() {
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

	projectRepository := repositories.NewProjectRepository(db)
	suite.projectService = services.NewProjectService(projectRepository)
}

func (suite *ProjectServiceTestSuite) TestCreateProject() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Die Hard",
		UserId:         "123123",
		OrganizationId: "123123",
		KeyPrefix:      "DH",
		Type:           "global",
	})

	suite.Nil(err)
}

func (suite *ProjectServiceTestSuite) TestCreateProjectWithInvalidUser() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Invalid User",
		UserId:         "INoExist",
		OrganizationId: "123123",
		KeyPrefix:      "IU",
		Type:           "global",
	})

	suite.NotNil(err)
}

func (suite *ProjectServiceTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestProjectTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectServiceTestSuite))
}
