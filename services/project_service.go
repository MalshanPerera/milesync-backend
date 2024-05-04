package services

import (
	"context"
	"errors"
	errpkg "jira-for-peasants/errors"
	repo "jira-for-peasants/repositories"
	"strings"
)

type ProjectService struct {
	projectRepository *repo.ProjectRepository
}

type CreateProjectParams struct {
	UserId         string
	OrganizationId string
	Name           string
	KeyPrefix      string
	Type           string
}

func NewProjectService(projectRepo *repo.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepo,
	}
}

func createKeyPrefix(key string) string {
	return strings.ToUpper(key)
}

func (s *ProjectService) CreateProject(ctx context.Context, params CreateProjectParams) (repo.ProjectModel, error) {
	_, err := s.projectRepository.GetProjectByKeyPrefix(ctx, params.KeyPrefix)

	if !errors.Is(err, errpkg.NoResults) {
		return repo.ProjectModel{}, errors.New(errpkg.ProjectExists)
	}

	project, err := s.projectRepository.CreateProject(ctx, repo.CreateProjectParams{
		UserId:         params.UserId,
		OrganizationId: params.OrganizationId,
		Name:           params.Name,
		KeyPrefix:      createKeyPrefix(params.KeyPrefix),
		Type:           params.Type,
	})

	if err != nil {
		return repo.ProjectModel{}, err
	}

	return project, nil
}

func (s *ProjectService) GetProjectByKeyPrefix(ctx context.Context, keyPrefix string) (repo.ProjectModel, error) {

	project, err := s.projectRepository.GetProjectByKeyPrefix(ctx, keyPrefix)

	if err != nil {
		return repo.ProjectModel{}, err
	}

	return project, nil
}

func (s *ProjectService) GetProjectKeyPrefixUsed(ctx context.Context, keyPrefix string) (bool, error) {
	key := createKeyPrefix(keyPrefix)
	return s.projectRepository.GetProjectKeyPrefixUsed(ctx, key)
}
