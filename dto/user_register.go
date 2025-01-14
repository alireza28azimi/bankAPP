package dto

type UserInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
type RegisterRequest struct {
	PhoneNumber string
	Name        string
	Password    string
	Email       string
}
type RegisterResponse struct {
	User UserInfo `json:"user"`
}
