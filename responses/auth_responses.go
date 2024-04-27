package responses

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
}

type RegisterUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewLoginResponse(token, refreshToken string, exp int64) *LoginResponse {
	return &LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Exp:          exp,
	}
}

func NewRegisterUserResponse(id, name, email string) *RegisterUserResponse {
	return &RegisterUserResponse{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
