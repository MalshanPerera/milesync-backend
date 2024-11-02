package requests

type RegisterUserRequest struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required"`
	FirstName        string `json:"first_name" validate:"required"`
	LastName         string `json:"last_name" validate:"required"`
	OrganizationSlug string `json:"organization_slug" validate:"required"`
}

type AddUserToOrganizationRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
