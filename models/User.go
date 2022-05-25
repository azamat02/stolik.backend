package models

type User struct {
	Id       uint   `json:"id"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
