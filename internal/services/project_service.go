package services

import (
	"context"
	"errors"
	repo "jira-for-peasants/internal/repositories"
	errpkg "jira-for-peasants/pkg/errors"
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

type UpdateProjectParams struct {
	ID        string
	UserId    string
	Name      string
	KeyPrefix string
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

	if len(params.KeyPrefix) > 4 {
		return repo.ProjectModel{}, errors.New(errpkg.KeyPrefixTooLong)
	}

	project, err := s.projectRepository.CreateProject(ctx, repo.CreateProjectParams{
		UserID:         params.UserId,
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

func (s *ProjectService) UpdateProject(ctx context.Context, params UpdateProjectParams) (repo.ProjectModel, error) {
	_, err := s.projectRepository.GetProjectById(ctx, params.ID, params.UserId)

	if err != nil {
		return repo.ProjectModel{}, err
	}

	project, err := s.projectRepository.UpdateProject(ctx, repo.UpdateProjectParams{
		ID:        params.ID,
		UserID:    params.UserId,
		Name:      params.Name,
		KeyPrefix: createKeyPrefix(params.KeyPrefix),
	})

	if err != nil {
		return repo.ProjectModel{}, err
	}

	return project, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string, userId string) error {
	_, err := s.projectRepository.GetProjectById(ctx, id, userId)

	if err != nil {
		return err
	}

	return s.projectRepository.DeleteProject(ctx, id, userId)
}

func (s *ProjectService) GetProjects(ctx context.Context, userId string, organizationId string) ([]repo.ProjectModel, error) {
	return s.projectRepository.GetProjects(ctx, userId, organizationId)
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
