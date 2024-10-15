package services

import (
	"context"
	repo "jira-for-peasants/repositories"
)

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

type StatusService struct {
	statusRepository *repo.StatusRepository
}

func NewStatusService(statusRepo *repo.StatusRepository) *StatusService {
	return &StatusService{
		statusRepository: statusRepo,
	}
}

func (s *StatusService) CreateStatus(ctx context.Context, params CreateStatusParams) (repo.StatusModel, error) {
	tx, err := s.statusRepository.BeginTx(ctx)
	if err != nil {
		return repo.StatusModel{}, err
	}

	defer func() {
		err = s.statusRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	status, err := s.statusRepository.CreateStatus(ctx, tx, repo.CreateStatusParams{
		Name:           params.Name,
		Color:          params.Color,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if err != nil {
		return repo.StatusModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.StatusModel{}, err
	}

	return status, nil
}

func (s *StatusService) UpdateStatus(ctx context.Context, params UpdateStatusParams) (repo.StatusModel, error) {
	tx, err := s.statusRepository.BeginTx(ctx)
	if err != nil {
		return repo.StatusModel{}, err
	}

	defer func() {
		err = s.statusRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	status, err := s.statusRepository.UpdateStatus(ctx, tx, repo.UpdateStatusParams{
		ID:    params.ID,
		Name:  params.Name,
		Color: params.Color,
	})

	if err != nil {
		return repo.StatusModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.StatusModel{}, err
	}

	return status, nil
}

func (s *StatusService) DeleteStatus(ctx context.Context, params DeleteStatusParams) error {
	tx, err := s.statusRepository.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = s.statusRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	err = s.statusRepository.DeleteStatus(ctx, tx, repo.DeleteStatusParams{
		ID:             params.ID,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *StatusService) GetStatuses(ctx context.Context, params GetStatusesParams) ([]repo.StatusModel, error) {
	tx, err := s.statusRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = s.statusRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	statuses, err := s.statusRepository.GetStatuses(ctx, tx, repo.GetStatusesParams{
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return statuses, nil
}
