package requests

type CreateStatusRequest struct {
	Name           string `json:"name" validate:"required"`
	Color          string `json:"color" validate:"required"`
	ProjectId      string `json:"project_id" validate:"required"`
	OrganizationId string `json:"organization_id" validate:"required"`
}

type UpdateStatusRequest struct {
	ID             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Color          string `json:"color" validate:"required"`
	ProjectId      string `json:"project_id" validate:"required"`
	OrganizationId string `json:"organization_id" validate:"required"`
	CreatedAt      string `json:"created_at" validate:"required"`
	UpdatedAt      string `json:"updated_at" validate:"required"`
}

type DeleteStatusRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetStatusesRequest struct {
	ProjectId      string `json:"project_id" validate:"required"`
	OrganizationId string `json:"organization_id" validate:"required"`
}
