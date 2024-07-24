package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"

	"github.com/jackc/pgx/v5"
)

type RoleModel db.Role
type RoleProjectModel db.RolesProject

type RoleRepository struct {
	*datastore.Trx
	db *datastore.DB
}

type CreateRoleParams struct {
	Name           string
	OrganizationId string
	DefaultId      *string
	Description    string
	Permissions    []string
}

type UpdateRoleParams struct {
	Name           string
	OrganizationId string
	Description    string
}

type RolePermissionParams struct {
	OrganizationId string
	RoleId         string
	Permission     string
}

type RoleUserParams struct {
	UserId    string
	RoleId    string
	ProjectId string
}

func NewRoleRepository(db *datastore.DB) *RoleRepository {
	return &RoleRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (repo *RoleRepository) CreateNewRole(ctx context.Context, tx pgx.Tx, params CreateRoleParams) (RoleModel, error) {
	role, err := repo.db.GetQuery().WithTx(tx).CreateRole(context.Background(), db.CreateRoleParams{
		OrganizationID: params.OrganizationId,
		DefaultID:      params.DefaultId,
		Name:           params.Name,
		Description:    params.Description,
		Permissions:    params.Permissions,
	})

	if err != nil {
		return RoleModel{}, err
	}

	return RoleModel(role), nil
}

func (repo *RoleRepository) DeleteRole(ctx context.Context, roleId string) error {
	err := repo.db.GetQuery().DeleteRole(ctx, roleId)

	if err != nil {
		return err
	}

	return nil
}

func (repo *RoleRepository) GetRoleById(ctx context.Context, roleId string) (RoleModel, error) {
	res, err := repo.db.GetQuery().GetRole(ctx, roleId)
	return RoleModel(res), err
}

func (repo *RoleRepository) UpdateRole(ctx context.Context, tx pgx.Tx, roleId string, params UpdateRoleParams) (RoleModel, error) {
	role, err := repo.db.GetQuery().WithTx(tx).UpdateRole(ctx, db.UpdateRoleParams{
		ID:             roleId,
		OrganizationID: params.OrganizationId,
		Name:           params.Name,
		Description:    params.Description,
	})

	if err != nil {
		return RoleModel{}, err
	}

	return RoleModel(role), nil
}

func (repo *RoleRepository) AddPermissionToRole(ctx context.Context, tx pgx.Tx, params RolePermissionParams) (RoleModel, error) {
	role, err := repo.db.GetQuery().WithTx(tx).AddPermissionToRole(ctx, db.AddPermissionToRoleParams{
		OrganizationID: params.OrganizationId,
		ID:             params.RoleId,
		ArrayAppend:    params.Permission,
	})

	if err != nil {
		return RoleModel{}, err
	}

	return RoleModel(role), nil
}

func (repo *RoleRepository) RemovePermissionFromRole(ctx context.Context, tx pgx.Tx, params RolePermissionParams) error {
	err := repo.db.GetQuery().WithTx(tx).RemovePermissionFromRole(ctx, db.RemovePermissionFromRoleParams{
		OrganizationID: params.OrganizationId,
		ID:             params.RoleId,
		ArrayRemove:    params.Permission,
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo *RoleRepository) AddRoleToUser(ctx context.Context, tx pgx.Tx, params RoleUserParams) (RoleProjectModel, error) {
	role, err := repo.db.GetQuery().WithTx(tx).AddRoleToUser(ctx, db.AddRoleToUserParams{
		UserID:    params.UserId,
		RoleID:    params.RoleId,
		ProjectID: params.ProjectId,
	})

	if err != nil {
		return RoleProjectModel{}, err
	}

	return RoleProjectModel(role), nil
}

func (repo *RoleRepository) RemoveRoleFromUser(ctx context.Context, tx pgx.Tx, userId string, roleId string, projectId string) error {
	err := repo.db.GetQuery().WithTx(tx).RemoveRoleFromUser(ctx, db.RemoveRoleFromUserParams{
		UserID:    userId,
		RoleID:    roleId,
		ProjectID: projectId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo *RoleRepository) GetRolesForOrganization(ctx context.Context, projectId string) ([]RoleModel, error) {
	roles, err := repo.db.GetQuery().GetRolesForOrganization(ctx, projectId)
	if err != nil {
		return nil, err
	}

	var result []RoleModel
	for _, role := range roles {
		result = append(result, RoleModel(role))
	}

	return result, nil
}

func (repo *RoleRepository) GetRolesForUser(ctx context.Context, userId string, organizationId string) ([]RoleModel, error) {
	roles, err := repo.db.GetQuery().GetRolesForUser(ctx, db.GetRolesForUserParams{
		UserID:         userId,
		OrganizationID: organizationId,
	})

	if err != nil {
		return nil, err
	}

	var result []RoleModel
	for _, role := range roles {
		result = append(result, RoleModel(role))
	}

	return result, nil
}
