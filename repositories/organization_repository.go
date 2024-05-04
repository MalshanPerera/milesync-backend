package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"
)

type OrganizationModel db.Organization

type OrganizationRepository struct {
	*datastore.Trx
	db *datastore.DB
}

type CreateOrganizationParams struct {
	Name   string
	UserId string
	Slug   string
}

func NewOrganizationRepository(db *datastore.DB) *OrganizationRepository {
	return &OrganizationRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (s *OrganizationRepository) CreateOrganization(ctx context.Context, params CreateOrganizationParams) (OrganizationModel, error) {
	organization, err := s.db.GetQuery().CreateOrganization(ctx, db.CreateOrganizationParams{
		Name:   params.Name,
		UserID: params.UserId,
		Slug:   params.Slug,
	})

	if err != nil {
		return OrganizationModel{}, err
	}

	return OrganizationModel(organization), nil
}

func (s *OrganizationRepository) DeleteOrganization(ctx context.Context, userId string) error {

	err := s.db.GetQuery().DeleteOrganization(ctx, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrganizationRepository) GetOrganizationSlugUsed(ctx context.Context, slug string) (bool, error) {
	return s.db.GetQuery().GetOrganizationSlugUsed(ctx, slug)
}

func (s *OrganizationRepository) GetOrganization(ctx context.Context, slug string) (OrganizationModel, error) {
	res, err := s.db.GetQuery().GetOrganization(ctx, slug)
	return OrganizationModel(res), err
}

func (s *OrganizationRepository) GetOrganizationByUserId(ctx context.Context, userId string) (OrganizationModel, error) {
	res, err := s.db.GetQuery().GetOrganizationByUserId(ctx, userId)
	if err != nil {
		return OrganizationModel{}, err
	}

	return OrganizationModel(res), nil
}
