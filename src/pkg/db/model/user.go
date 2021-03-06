package model

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password []byte `json:"password" validate:"required"`
	Hobby    string `json:"hobby"`
}
