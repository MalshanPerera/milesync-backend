package services

import (
	"context"
	repo "jira-for-peasants/internal/repositories"
)

// Get all permissions
// Add role to project (complicated)

type RoleService struct {
	roleRepository *repo.RoleRepository
}

func NewRoleService(roleRepository *repo.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

type CreateRoleParams struct {
	Name           string
	OrganizationId string
	Description    string
}

type UpdateRoleParams struct {
	Name           string
	OrganizationId string
	Description    string
}

type AddPermissionToRoleParams struct {
	RoleId     string
	Permission string
}

type RemovePermissionFromRoleParams struct {
	RoleId     string
	Permission string
}

func (s *RoleService) CreateRole(ctx context.Context, params CreateRoleParams) (repo.RoleModel, error) {
	tx, err := s.roleRepository.BeginTx(ctx)

	if err != nil {
		return repo.RoleModel{}, err
	}
	defer func() {
		err = s.roleRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	role, err := s.roleRepository.CreateNewRole(ctx, tx, repo.CreateRoleParams{
		Name:           params.Name,
		OrganizationId: params.OrganizationId,
		Description:    params.Description,
		Permissions:    []string{},
	})

	if err != nil {
		return repo.RoleModel{}, err
	}

	err = s.roleRepository.CommitTx(ctx, tx)

	if err != nil {
		return repo.RoleModel{}, err
	}

	return role, nil
}

func (s *RoleService) DeleteRole(ctx context.Context, roleId string) error {

	role, err := s.roleRepository.GetRoleById(ctx, roleId)

	if err != nil {
		return err
	}

	err = s.roleRepository.DeleteRole(ctx, role.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *RoleService) UpdateRole(ctx context.Context, roleId string, params UpdateRoleParams) (repo.RoleModel, error) {
	tx, err := s.roleRepository.BeginTx(ctx)

	if err != nil {
		return repo.RoleModel{}, err
	}
	defer func() {
		err = s.roleRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	role, err := s.roleRepository.UpdateRole(ctx, tx, roleId, repo.UpdateRoleParams{
		Name:           params.Name,
		OrganizationId: params.OrganizationId,
		Description:    params.Description,
	})

	if err != nil {
		return repo.RoleModel{}, err
	}

	err = s.roleRepository.CommitTx(ctx, tx)

	if err != nil {
		return repo.RoleModel{}, err
	}

	return role, nil
}

func (s *RoleService) AddPermissionToRole(ctx context.Context, organizationId string, params AddPermissionToRoleParams) (repo.RoleModel, error) {
	tx, err := s.roleRepository.BeginTx(ctx)

	if err != nil {
		return repo.RoleModel{}, err
	}
	defer func() {
		err = s.roleRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	role, err := s.roleRepository.AddPermissionToRole(ctx, tx, repo.RolePermissionParams{
		RoleId:         params.RoleId,
		OrganizationId: organizationId,
		Permission:     params.Permission,
	})

	if err != nil {
		return repo.RoleModel{}, err
	}

	err = s.roleRepository.CommitTx(ctx, tx)

	if err != nil {
		return repo.RoleModel{}, err
	}

	return role, nil
}

func (s *RoleService) RemovePermissionFromRole(ctx context.Context, organizationId string, params RemovePermissionFromRoleParams) error {
	tx, err := s.roleRepository.BeginTx(ctx)

	if err != nil {
		return err
	}
	defer func() {
		err = s.roleRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	err = s.roleRepository.RemovePermissionFromRole(ctx, tx, repo.RolePermissionParams{
		RoleId:         params.RoleId,
		OrganizationId: organizationId,
		Permission:     params.Permission,
	})

	if err != nil {
		return err
	}

	err = s.roleRepository.CommitTx(ctx, tx)

	if err != nil {
		return err
	}

	return nil
}

func (s *RoleService) GetRoleById(ctx context.Context, roleId string) (repo.RoleModel, error) {
	return s.roleRepository.GetRoleById(ctx, roleId)
}

func (s *RoleService) GetAllRoleForOrganization(ctx context.Context, organizationId string) ([]repo.RoleModel, error) {
	return s.roleRepository.GetRolesForOrganization(ctx, organizationId)
}

func (s *RoleService) GetAllRoleForUser(ctx context.Context, userId string, organizationId string) ([]repo.RoleModel, error) {
	return s.roleRepository.GetRolesForUser(ctx, userId, organizationId)
}

func (s *RoleService) AddRoleToUser(ctx context.Context, userId string, roleId string, projectId string) (repo.RoleProjectModel, error) {
	tx, err := s.roleRepository.BeginTx(ctx)

	if err != nil {
		return repo.RoleProjectModel{}, err
	}
	defer func() {
		err = s.roleRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	role, err := s.roleRepository.AddRoleToUser(ctx, tx, repo.RoleUserParams{
		UserId:    userId,
		RoleId:    roleId,
		ProjectId: projectId,
	})

	if err != nil {
		return repo.RoleProjectModel{}, err
	}

	err = s.roleRepository.CommitTx(ctx, tx)

	if err != nil {
		return repo.RoleProjectModel{}, err
	}

	return role, nil
}

func (s *RoleService) RemoveRoleFromUser(ctx context.Context, userId string, roleId string, projectId string) error {
	tx, err := s.roleRepository.BeginTx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		err = s.roleRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	err = s.roleRepository.RemoveRoleFromUser(ctx, tx, userId, roleId, projectId)

	if err != nil {
		return err
	}

	err = s.roleRepository.CommitTx(ctx, tx)

	if err != nil {
		return err
	}

	return nil
}
