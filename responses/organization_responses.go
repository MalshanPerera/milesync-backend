package responses

type OrganizationResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedBy string `json:"created_by"`
}

func NewOrganizationResponse(
	ID string,
	Name string,
	Slug string,
	CreatedBy string,
) *OrganizationResponse {
	return &OrganizationResponse{
		ID:        ID,
		Name:      Name,
		Slug:      Slug,
		CreatedBy: CreatedBy,
	}
}
