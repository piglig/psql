package components

import (
	"fmt"
	"github.com/stoewer/go-strcase"
	"psql/pkg/psql"
	"psql/pkg/utils"
	"strings"
)

// GenerateDaoObjFun dao 对象生成函数
/**
example:
	func NewGame() *GameDao {
		return &GameDao{}
	}
*/
func GenerateDaoObjFun(model string) string {
	res := fmt.Sprintf("func New%sDao() *%sDao {\n", model, model)
	res += fmt.Sprintf("\treturn &%sDao{}\n", model)
	res += fmt.Sprintf("}\n\n")
	return res
}

// GenerateDaoCreateFunc dao C 增
/**
example:
	func (e *ErrorCode) CreateErrorCode(ctx *app.Context, errCode model.ErrorCode) error {
		rst := ctx.DB.Omit("create_time").Create(errCode)
		if rst.Error != nil {
			return rst.Error
		}
		return nil
	}
*/
func GenerateDaoCreateFunc(model string) string {
	res := fmt.Sprintf("func (*%sDao) Create%s(ctx *app.Context, %s model.%s) error {\n", model, model,
		strings.ToLower(model), model)
	res += fmt.Sprintf("\trst := ctx.DB.Create(&%s)\n", strings.ToLower(model))
	res += fmt.Sprintf("\tif rst.Error != nil {\n")
	res += fmt.Sprintf("\t\treturn rst.Error\n")
	res += fmt.Sprintf("}\n")
	res += fmt.Sprintf("\treturn nil\n")
	res += fmt.Sprintf("}\n\n")
	return res
}

// GenerateDaoDeleteByKeyFunc dao D 单一字段删
/**
example:
	func (g *GameDao) DeleteGameByGameId(ctx *app.Context, gameId string) bool {
		rst := ctx.DB.Where("game_id = ?", gameId).Debug().Delete(&model.Game{})
		if rst.Error != nil {
			return false
		}

		return true
	}
*/
func GenerateDaoDeleteByKeyFunc(model, key, keyType string) string {
	lowerKey := strings.ToLower(key)
	res := fmt.Sprintf("func (*%sDao) Remove%sBy%s(ctx *app.Context, %s %s ) error {\n", model, model,
		utils.FirstLetterToUpper(key), key, keyType)
	res += fmt.Sprintf("\trst := ctx.DB.Where(\"%s = ?\", %s).Delete(&model.%s{})\n", lowerKey, lowerKey, model)
	res += fmt.Sprintf("\tif rst.Error != nil {\n")
	res += fmt.Sprintf("\t\treturn rst.Error\n")
	res += fmt.Sprintf("}\n")
	res += fmt.Sprintf("\treturn nil\n")
	res += fmt.Sprintf("}\n\n")
	return res
}

// GenerateDaoDeleteByKeysFunc dao D 多字段删
/**
example:
	func (g *GameDao) DeleteGameByGameIdAndGameType(ctx *app.Context, gameId string) bool {
		rst := ctx.DB.Where("game_id = ?", gameId).Debug().Delete(&model.Game{})
		if rst.Error != nil {
			return false
		}

		return true
	}
*/
func GenerateDaoDeleteByKeysFunc(model string, keys []string) string {
	funcName := getDaoDeleteKeysStr(keys)
	funcParamsStr := ""
	for i, key := range keys {
		if i == 0 {
			funcParamsStr += getDaoDeleteKeyPairStr(key)
		} else {
			funcParamsStr += ", " + getDaoDeleteKeyPairStr(key)
		}
	}

	res := fmt.Sprintf("func (*%sDao) Remove%sBy%s(ctx *app.Context, %s) error {\n", model, model,
		funcName, funcParamsStr)
	res += fmt.Sprintf("\trst := ctx.DB.Where(%s).Delete(&model.%s{})\n", getDaoWherePairStr(keys), model)
	res += fmt.Sprintf("\tif rst.Error != nil {\n")
	res += fmt.Sprintf("\t\treturn rst.Error\n")
	res += fmt.Sprintf("\t}\n")
	res += fmt.Sprintf("\treturn nil\n")
	res += fmt.Sprintf("}\n\n")
	return res
}

func getDaoDeleteKeysStr(keys []string) string {
	res := ""
	for i, key := range keys {
		if i == 0 {
			res += utils.FirstLetterToUpper(strcase.LowerCamelCase(key))
		} else {
			res += strcase.UpperCamelCase(key)
		}

	}

	return res
}

func getDaoDeleteKeyPairStr(key string) string {
	res := strcase.LowerCamelCase(key) + " " + psql.GetFieldType(key)
	return res
}

func getDaoWherePairStr(keys []string) string {
	res := "\""
	var tempKeys []string
	for i, key := range keys {
		tempKeys = append(tempKeys, strcase.LowerCamelCase(key))
		if i == 0 {
			res += strcase.SnakeCase(key) + " = ?"
		} else {
			res += " and " + strcase.SnakeCase(key) + " = ?"
		}
	}
	res += "\", " + utils.SliceToStrByDelimiter(tempKeys, ", ")
	return res
}
