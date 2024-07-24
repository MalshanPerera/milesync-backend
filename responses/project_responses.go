package responses

type ProjectResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	KeyPrefix      string `json:"key_prefix"`
	Type           string `json:"type"`
	OrganizationId string `json:"organization_id"`
	CreatedBy      string `json:"created_by"`
}

// func NewProjectResponse(
// 	ID string,
// 	Name string,
// 	KeyPrefix string,
// 	Type string,
// 	OrganizationId string,
// 	CreatedBy string,
// ) *ProjectResponse {
// 	return &ProjectResponse{
// 		ID:             ID,
// 		Name:           Name,
// 		KeyPrefix:      KeyPrefix,
// 		Type:           Type,
// 		OrganizationId: OrganizationId,
// 		CreatedBy:      CreatedBy,
// 	}
// }
