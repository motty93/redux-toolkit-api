package model

type User struct {
	Base
	Name     string `json:"name"`
	Password string `json:"password"`
	Hobby    string `json:"hobby"`
}
