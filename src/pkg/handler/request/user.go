package request

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validte:"required,password"`
}
