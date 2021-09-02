package model

type User struct {
	Gender int
	Name   string
	Email  string
}

func (*User) TableName() string {
	return "user"
}
