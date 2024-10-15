package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type LabelModel db.Label

type CreateLabelParams struct {
	Name           string
	Color          string
	ProjectID      string
	OrganizationID string
}

type UpdateLabelParams struct {
	ID    string
	Name  string
	Color string
}

type DeleteLabelParams struct {
	ID             string
	ProjectID      string
	OrganizationID string
}

type GetLabelsParams struct {
	ProjectID      string
	OrganizationID string
}

type LabelRepository struct {
	*datastore.Trx
	db *datastore.DB
}

func NewLabelRepository(db *datastore.DB) *LabelRepository {
	return &LabelRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (repo *LabelRepository) CreateLabel(ctx context.Context, tx pgx.Tx, params CreateLabelParams) (LabelModel, error) {
	label, e := repo.db.GetQuery().WithTx(tx).CreateLabel(ctx, db.CreateLabelParams{
		Name:           params.Name,
		Color:          params.Color,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if e != nil {
		return LabelModel{}, e
	}

	return LabelModel(label), nil
}

func (repo *LabelRepository) UpdateLabel(ctx context.Context, tx pgx.Tx, params UpdateLabelParams) (LabelModel, error) {
	label, e := repo.db.GetQuery().WithTx(tx).UpdateLabel(ctx, db.UpdateLabelParams{
		ID:    params.ID,
		Name:  params.Name,
		Color: params.Color,
	})

	if e != nil {
		return LabelModel{}, e
	}

	return LabelModel(label), nil
}

func (repo *LabelRepository) DeleteLabel(ctx context.Context, tx pgx.Tx, params DeleteLabelParams) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteLabel(ctx, db.DeleteLabelParams{
		ID:             params.ID,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if e != nil {
		return e
	}

	return nil
}

func (repo *LabelRepository) GetLabels(ctx context.Context, tx pgx.Tx, params GetLabelsParams) ([]LabelModel, error) {
	labels, e := repo.db.GetQuery().WithTx(tx).GetLabels(ctx, db.GetLabelsParams{
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if e != nil {
		return []LabelModel{}, e
	}

	results := make([]LabelModel, 0)
	for _, label := range labels {
		results = append(results, LabelModel(label))
	}

	return results, nil

}
