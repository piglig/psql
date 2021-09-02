package main

import (
	"psql/pkg/file"
	"psql/pkg/logger"
	"psql/pkg/psql"
)

/**
TODO
1. 数据表名驼峰和下划线转换 bug
*/

func main() {
	logger.InitLogger()
	psql.InitFile("./user.json")
	file.InitOutFileConfig(".", "user.go")

	file.Run()
}
