package dto

type LoginRequest struct {
	PhoneNumber string
	Password    string
}
type LoginResponse struct {
	User UserInfo `json:"user"`
}
