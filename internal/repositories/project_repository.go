package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"
)

type ProjectModel db.Project

type ProjectRepository struct {
	*datastore.Trx
	db *datastore.DB
}

type CreateProjectParams struct {
	OrganizationId string
	UserID         string
	Name           string
	KeyPrefix      string
	Type           string
}

type UpdateProjectParams struct {
	ID        string
	UserID    string
	Name      string
	KeyPrefix string
}

func NewProjectRepository(db *datastore.DB) *ProjectRepository {
	return &ProjectRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (s *ProjectRepository) CreateProject(ctx context.Context, params CreateProjectParams) (ProjectModel, error) {
	project, err := s.db.GetQuery().CreateProject(ctx, db.CreateProjectParams{
		OrganizationID: params.OrganizationId,
		UserID:         params.UserID,
		Name:           params.Name,
		KeyPrefix:      params.KeyPrefix,
		Type:           params.Type,
	})

	if err != nil {
		return ProjectModel{}, err
	}

	return ProjectModel(project), nil

}

func (s *ProjectRepository) UpdateProject(ctx context.Context, params UpdateProjectParams) (ProjectModel, error) {
	project, err := s.db.GetQuery().UpdateProject(ctx, db.UpdateProjectParams{
		ID:        params.ID,
		UserID:    params.UserID,
		Name:      params.Name,
		KeyPrefix: params.KeyPrefix,
	})

	if err != nil {
		return ProjectModel{}, err
	}

	return ProjectModel(project), nil
}

func (s *ProjectRepository) DeleteProject(ctx context.Context, id string, userId string) error {
	return s.db.GetQuery().DeleteProject(ctx, db.DeleteProjectParams{
		ID:     id,
		UserID: userId,
	})
}

func (s *ProjectRepository) GetProjectByKeyPrefix(ctx context.Context, keyPrefix string) (ProjectModel, error) {
	project, err := s.db.GetQuery().GetProjectByKeyPrefix(ctx, keyPrefix)

	if err != nil {
		return ProjectModel{}, err
	}

	return ProjectModel(project), nil
}

func (s *ProjectRepository) GetProjectKeyPrefixUsed(ctx context.Context, keyPrefix string) (bool, error) {
	return s.db.GetQuery().GetProjectKeyPrefixUsed(ctx, keyPrefix)
}

func (s *ProjectRepository) GetProjectById(ctx context.Context, id string, userId string) (ProjectModel, error) {
	project, err := s.db.GetQuery().GetProject(ctx, db.GetProjectParams{
		ID:     id,
		UserID: userId,
	})

	if err != nil {
		return ProjectModel{}, err
	}

	return ProjectModel(project), nil
}

func (s *ProjectRepository) GetProjects(ctx context.Context, userId string, organizationId string) ([]ProjectModel, error) {
	projects, err := s.db.GetQuery().GetProjects(ctx, db.GetProjectsParams{
		UserID:         userId,
		OrganizationID: organizationId,
	})

	if err != nil {
		return nil, err
	}

	var projectModels []ProjectModel
	for _, project := range projects {
		projectModels = append(projectModels, ProjectModel(project))
	}

	return projectModels, nil
}
