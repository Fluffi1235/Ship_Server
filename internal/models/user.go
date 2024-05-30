package models

type User struct {
	Id           int64
	UserName     string
	PasswordHash string
	Name         string
	Surname      string
	Department   string
	Rank         string
	Email        string
	PhotoUrl     string
}
