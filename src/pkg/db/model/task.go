package model

type Task struct {
	Base
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}
