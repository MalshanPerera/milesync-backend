package requests

type CreateProjectRequest struct {
	OrganizationId string `json:"organization_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	KeyPrefix      string `json:"key_prefix" validate:"required,max=4"`
	Type           string `json:"type" validate:"required,oneof=global project department"`
}
