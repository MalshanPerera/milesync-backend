package services_test

import (
	"context"
	datastore "jira-for-peasants/db"
	"jira-for-peasants/internal/repositories"
	services "jira-for-peasants/internal/services"
	errpkg "jira-for-peasants/pkg/errors"
	testUtils "jira-for-peasants/test/test_utils"
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
		"test-organizations.sql",
		"test-projects.sql",
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
		OrganizationId: "151515",
		KeyPrefix:      "DH",
		Type:           "global",
	})

	suite.Nil(err)
}

func (suite *ProjectServiceTestSuite) TestCreateProjectWithInvalidUser() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Invalid User",
		UserId:         "INoExist",
		OrganizationId: "151515",
		KeyPrefix:      "IU",
		Type:           "global",
	})

	suite.NotNil(err)
}

func (suite *ProjectServiceTestSuite) TestCreateProjectWithInvalidOrganization() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Invalid Organization",
		UserId:         "123123",
		OrganizationId: "INoExist",
		KeyPrefix:      "IO",
		Type:           "global",
	})

	suite.NotNil(err)
}

func (suite *ProjectServiceTestSuite) TestCreateProjectWithInvalidType() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Invalid Type",
		UserId:         "123123",
		OrganizationId: "151515",
		KeyPrefix:      "IT",
		Type:           "invalid",
	})

	suite.NotNil(err)
}

func (suite *ProjectServiceTestSuite) TestCreateProjectWithMoreThan4KeyPrefix() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Invalid Key Prefix",
		UserId:         "123123",
		OrganizationId: "151515",
		KeyPrefix:      "InvalidKeyPrefix",
		Type:           "global",
	})

	suite.EqualError(err, errpkg.KeyPrefixTooLong)
}

func (suite *ProjectServiceTestSuite) TestCreateProjectWithExistingKeyPrefix() {
	_, err := suite.projectService.CreateProject(suite.ctx, services.CreateProjectParams{
		Name:           "Die Hard",
		UserId:         "123123",
		OrganizationId: "151515",
		KeyPrefix:      "P1",
		Type:           "global",
	})

	suite.EqualError(err, errpkg.ProjectExists)
}

func (suite *ProjectServiceTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestProjectTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectServiceTestSuite))
}
