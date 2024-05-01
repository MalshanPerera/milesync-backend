package services

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"
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

	organization, e := s.db.GetQuery().CreateOrganization(ctx, db.CreateOrganizationParams{
		Name:   params.Name,
		UserID: params.UserId,
		Slug:   createSlug(params.Name),
	})

	if e != nil {
		return db.Organization{}, e
	}

	return organization, nil
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, userId string) error {

	e := s.db.GetQuery().DeleteOrganization(ctx, userId)

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
