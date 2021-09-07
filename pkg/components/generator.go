package components

import (
	"fmt"
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
	res := fmt.Sprintf("func (*%sDao) Delete%sBy%s(ctx *app.Context, %s %s ) error {\n", model, model,
		key, key, keyType)
	res += fmt.Sprintf("\trst := ctx.DB.Where(\"%s = ?\", %s)\n", lowerKey, lowerKey)
	res += fmt.Sprintf("\tif rst.Error != nil {\n")
	res += fmt.Sprintf("\t\treturn rst.Error\n")
	res += fmt.Sprintf("}\n")
	res += fmt.Sprintf("\treturn nil\n")
	res += fmt.Sprintf("}\n\n")
	return res
}
