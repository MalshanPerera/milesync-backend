package services

import (
	"context"
	"jira-for-peasents/common"
	datastore "jira-for-peasents/db"
	db "jira-for-peasents/db/sqlc"
	"strings"
)

type OrganizationService struct {
	db *datastore.DB
}

func NewOrganizationService(db *datastore.DB) *OrganizationService {
	return &OrganizationService{
		db: db,
	}
}

type CreateOrganizationParams struct {
	Name   string
	UserId string
}

func createSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, params CreateOrganizationParams) (db.Organization, error) {
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return db.Organization{}, common.NewDBError(err.Error())
	}

	defer func() {
		err = s.db.RollbackTx(ctx, tx)
	}()

	organization, e := s.db.GetQuery().WithTx(tx).CreateOrganization(ctx, db.CreateOrganizationParams{
		Name:   params.Name,
		UserID: params.UserId,
		Slug:   createSlug(params.Name),
	})

	if e != nil {
		return db.Organization{}, e
	}

	e = s.db.CommitTx(ctx, tx)

	if e != nil {
		return db.Organization{}, e
	}

	return organization, nil
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, userId string) error {
	tx, err := s.db.BeginTx(ctx)

	if err != nil {
		return common.NewDBError(err.Error())
	}

	defer func() {
		err = s.db.RollbackTx(ctx, tx)
	}()

	e := s.db.GetQuery().WithTx(tx).DeleteOrganization(ctx, userId)

	if e != nil {
		return e
	}

	e = s.db.CommitTx(ctx, tx)

	if e != nil {
		return e
	}

	return nil
}

func (s *OrganizationService) GetOrganizationSlugUsed(ctx context.Context, name string) (bool, error) {
	slug := createSlug(name)
	return s.db.GetQuery().GetOrganizationSlugUsed(ctx, slug)
}

func (s *OrganizationService) GetOrganization(ctx context.Context, slug string) (db.Organization, error) {
	return s.db.GetQuery().GetOrganization(ctx, slug)
}
