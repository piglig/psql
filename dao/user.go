package dao

import (
	"psql/app"
	"psql/model"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (*UserDao) DeleteUserByname(ctx *app.Context, name string) error {
	rst := ctx.DB.Where("name = ?", name)
	if rst.Error != nil {
		return rst.Error
	}
	return nil
}

func (*UserDao) DeleteUserByemail(ctx *app.Context, email string) error {
	rst := ctx.DB.Where("email = ?", email)
	if rst.Error != nil {
		return rst.Error
	}
	return nil
}

func (*UserDao) CreateUser(ctx *app.Context, user model.User) error {
	rst := ctx.DB.Create(&user)
	if rst.Error != nil {
		return rst.Error
	}
	return nil
}
