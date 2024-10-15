package requests

type CreateLabelRequest struct {
	Name      string `json:"name" validate:"required"`
	Color     string `json:"color" validate:"required"`
	ProjectId string `json:"project_id" validate:"required"`
}

type UpdateLabelRequest struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Color string `json:"color" validate:"required"`
}

type DeleteLabelRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetLabelsRequest struct {
	ProjectId string `json:"project_id" validate:"required"`
}
