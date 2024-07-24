package requests

type CreateProjectRequest struct {
	Name      string `json:"name" validate:"required"`
	KeyPrefix string `json:"key_prefix" validate:"required,max=4"`
	Type      string `json:"type" validate:"required,oneof=global project department"`
}

type UpdateProjectRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	KeyPrefix string `json:"key_prefix" validate:"required,max=4"`
}
