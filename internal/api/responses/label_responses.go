package responses

type LabelResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	ProjectID string `json:"project_id"`
}
