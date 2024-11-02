package responses

type RoleResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type RoleListResponse []RoleResponse

func NewRoleResponse(
	ID string,
	Name string,
	Description string,
	Permissions []string,
) *RoleResponse {
	return &RoleResponse{
		ID:          ID,
		Name:        Name,
		Description: Description,
		Permissions: Permissions,
	}
}

func NewRoleListResponse(roles []RoleResponse) *RoleListResponse {
	return (*RoleListResponse)(&roles)
}
