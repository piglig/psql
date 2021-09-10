package dao

import (
	"psql/app"
	"psql/model"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (*UserDao) RemoveUserByName(ctx *app.Context, name string) error {
	rst := ctx.DB.Where("name = ?", name).Delete(&model.User{})
	if rst.Error != nil {
		return rst.Error
	}
	return nil
}

func (*UserDao) RemoveUserByEmail(ctx *app.Context, email string) error {
	rst := ctx.DB.Where("email = ?", email).Delete(&model.User{})
	if rst.Error != nil {
		return rst.Error
	}
	return nil
}

func (*UserDao) RemoveUserByNameGender(ctx *app.Context, name string, gender int) error {
	rst := ctx.DB.Where("name = ? and gender = ?", name, gender).Delete(&model.User{})
	if rst.Error != nil {
		return rst.Error
	}
	return nil
}

func (*UserDao) RemoveUserByNameEmail(ctx *app.Context, name string, email string) error {
	rst := ctx.DB.Where("name = ? and email = ?", name, email).Delete(&model.User{})
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
