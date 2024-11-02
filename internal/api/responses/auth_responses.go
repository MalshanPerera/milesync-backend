package responses

type AuthResponse struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Exp          int64  `json:"exp"`
}

func NewAuthResponse(
	ID string,
	FirstName string,
	LastName string,
	Email string,
	AccessToken string,
	RefreshToken string,
	Exp int64,
) *AuthResponse {
	return &AuthResponse{
		ID:           ID,
		FirstName:    FirstName,
		LastName:     LastName,
		Email:        Email,
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
		Exp:          Exp,
	}
}
