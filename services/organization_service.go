package services

import (
	"context"
	errpkg "jira-for-peasants/errors"
	repo "jira-for-peasants/repositories"
	"strings"
)

type OrganizationService struct {
	organizationRepository *repo.OrganizationRepository
}

func NewOrganizationService(organizationRepo *repo.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		organizationRepository: organizationRepo,
	}
}

type CreateOrganizationParams struct {
	Name   string
	UserId string
}

func createSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, params CreateOrganizationParams) (repo.OrganizationModel, error) {
	_, err := s.organizationRepository.GetOrganizationByUserId(ctx, params.UserId)

	if err != errpkg.NoResults {
		return repo.OrganizationModel{}, errpkg.ServiceError(errpkg.UserAlreadyHasOrganization)
	}

	organization, err := s.organizationRepository.CreateOrganization(ctx, repo.CreateOrganizationParams{
		Name:   params.Name,
		UserId: params.UserId,
		Slug:   createSlug(params.Name),
	})

	if err != nil {
		return repo.OrganizationModel{}, err
	}

	return organization, nil
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, userId string) error {

	err := s.organizationRepository.DeleteOrganization(ctx, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrganizationService) GetOrganizationSlugUsed(ctx context.Context, name string) (bool, error) {
	slug := createSlug(name)
	return s.organizationRepository.GetOrganizationSlugUsed(ctx, slug)
}

func (s *OrganizationService) GetOrganization(ctx context.Context, slug string) (repo.OrganizationModel, error) {
	return s.organizationRepository.GetOrganization(ctx, slug)
}
