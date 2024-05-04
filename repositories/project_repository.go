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
	UserId         string
	Name           string
	KeyPrefix      string
	Type           string
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
		UserID:         params.UserId,
		Name:           params.Name,
		KeyPrefix:      params.KeyPrefix,
		Type:           params.Type,
	})

	if err != nil {
		return ProjectModel{}, err
	}

	return ProjectModel(project), nil

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
