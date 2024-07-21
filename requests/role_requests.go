package requests

type CreateRoleRequest struct {
	Name           string `json:"name" validate:"required"`
	OrganizationId string `json:"organization_id" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

type UpdateRoleRequest struct {
	Name           string `json:"name" validate:"required"`
	OrganizationId string `json:"organization_id" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

type AddPermissionToRoleRequest struct {
	RoleId     string `json:"role_id" validate:"required"`
	Permission string `json:"permission" validate:"required"`
}

type RemovePermissionFromRoleRequest struct {
	RoleId     string `json:"role_id" validate:"required"`
	Permission string `json:"permission" validate:"required"`
}

type AddUserToRoleRequest struct {
	UserId    string `json:"user_id" validate:"required"`
	RoleId    string `json:"role_id" validate:"required"`
	ProjectId string `json:"project_id" validate:"required"`
}

type RemoveUserFromRoleRequest struct {
	UserId    string `json:"user_id" validate:"required"`
	RoleId    string `json:"role_id" validate:"required"`
	ProjectId string `json:"project_id" validate:"required"`
}
