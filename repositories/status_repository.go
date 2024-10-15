package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type StatusModel db.Status

type CreateStatusParams struct {
	Name           string
	Color          string
	ProjectID      string
	OrganizationID string
}

type UpdateStatusParams struct {
	ID    string
	Name  string
	Color string
}

type DeleteStatusParams struct {
	ID             string
	ProjectID      string
	OrganizationID string
}

type GetStatusesParams struct {
	ProjectID      string
	OrganizationID string
}

type StatusRepository struct {
	*datastore.Trx
	db *datastore.DB
}

func NewStatusRepository(db *datastore.DB) *StatusRepository {
	return &StatusRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (repo *StatusRepository) CreateStatus(ctx context.Context, tx pgx.Tx, params CreateStatusParams) (StatusModel, error) {
	status, e := repo.db.GetQuery().WithTx(tx).CreateStatus(ctx, db.CreateStatusParams{
		Name:           params.Name,
		Color:          params.Color,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if e != nil {
		return StatusModel{}, e
	}

	return StatusModel(status), nil
}

func (repo *StatusRepository) UpdateStatus(ctx context.Context, tx pgx.Tx, params UpdateStatusParams) (StatusModel, error) {
	status, e := repo.db.GetQuery().WithTx(tx).UpdateStatus(ctx, db.UpdateStatusParams{
		ID:    params.ID,
		Name:  params.Name,
		Color: params.Color,
	})

	if e != nil {
		return StatusModel{}, e
	}

	return StatusModel(status), nil
}

func (repo *StatusRepository) DeleteStatus(ctx context.Context, tx pgx.Tx, params DeleteStatusParams) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteStatus(ctx, db.DeleteStatusParams{
		ID:             params.ID,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if e != nil {
		return e
	}

	return nil
}

func (repo *StatusRepository) GetStatuses(ctx context.Context, tx pgx.Tx, params GetStatusesParams) ([]StatusModel, error) {
	statuses, e := repo.db.GetQuery().WithTx(tx).GetStatuses(ctx, db.GetStatusesParams{
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if e != nil {
		return []StatusModel{}, e
	}

	results := make([]StatusModel, 0)
	for _, status := range statuses {
		results = append(results, StatusModel(status))
	}

	return results, nil
}
