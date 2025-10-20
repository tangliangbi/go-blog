package types

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserLoginResponse struct {
	AceessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
