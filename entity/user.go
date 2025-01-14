package entity

type User struct {
	ID   uint
	Name string
	//password is always hashed
	Password    string
	PhoneNumber string
	Email       string
	Role        Role
}
