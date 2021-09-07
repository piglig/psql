package app

import "gorm.io/gorm"

type Context struct {
	DB *gorm.DB
}

var Ctx *Context
