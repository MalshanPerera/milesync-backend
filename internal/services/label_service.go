package services

import (
	"context"
	repo "jira-for-peasants/internal/repositories"
)

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

type LabelService struct {
	labelRepository *repo.LabelRepository
}

func NewLabelService(labelRepo *repo.LabelRepository) *LabelService {
	return &LabelService{
		labelRepository: labelRepo,
	}
}

func (s *LabelService) CreateLabel(ctx context.Context, params CreateLabelParams) (repo.LabelModel, error) {
	tx, err := s.labelRepository.BeginTx(ctx)
	if err != nil {
		return repo.LabelModel{}, err
	}

	defer func() {
		err = s.labelRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	label, err := s.labelRepository.CreateLabel(ctx, tx, repo.CreateLabelParams{
		Name:           params.Name,
		Color:          params.Color,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if err != nil {
		return repo.LabelModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.LabelModel{}, err
	}

	return label, nil
}

func (s *LabelService) UpdateLabel(ctx context.Context, params UpdateLabelParams) (repo.LabelModel, error) {
	tx, err := s.labelRepository.BeginTx(ctx)
	if err != nil {
		return repo.LabelModel{}, err
	}

	defer func() {
		err = s.labelRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	label, err := s.labelRepository.UpdateLabel(ctx, tx, repo.UpdateLabelParams{
		ID:    params.ID,
		Name:  params.Name,
		Color: params.Color,
	})

	if err != nil {
		return repo.LabelModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.LabelModel{}, err
	}

	return label, nil
}

func (s *LabelService) DeleteLabel(ctx context.Context, params DeleteLabelParams) error {
	tx, err := s.labelRepository.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = s.labelRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	err = s.labelRepository.DeleteLabel(ctx, tx, repo.DeleteLabelParams{
		ID:             params.ID,
		ProjectID:      params.ProjectID,
		OrganizationID: params.OrganizationID,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *LabelService) GetLabels(ctx context.Context, params GetLabelsParams) ([]repo.LabelModel, error) {
	tx, err := s.labelRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = s.labelRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	labels, err := s.labelRepository.GetLabels(ctx, tx, repo.GetLabelsParams{
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

	return labels, nil
}
