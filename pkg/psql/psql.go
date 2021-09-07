package psql

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"path/filepath"
	"psql/pkg/logger"
	"psql/pkg/utils"
)

const (
	INT64  = "int64"
	INT    = "int"
	STRING = "string"
)

var p = new(PSql)

type Model struct {
	StringAttributes []string `json:"string"`
	IntAttributes    []string `json:"int"`
}

type PJson struct {
	model Model `json:"model"`
}

type PSql struct {
	File
	pjson PJson `json:"model"`
}

type File struct {
	FileName string
	Content  string `json:"-"`
}

func InitFile(fp string) {
	f, err := os.Open(fp)
	defer f.Close()

	if err != nil {
		logger.Fatal("InitFile os.Open fail:%v", err)
	}

	fByte, err := io.ReadAll(f)
	if err != nil {
		logger.Fatal("InitFile io.ReadAll fail:%v", err)
	}

	p.Content = string(fByte)
	p.FileName = filepath.Base(fp)

	if !p.CheckFormat() {
		logger.Fatal("InitFile CheckFormat fail")
		return
	}

	logger.Info("InitFile success.")
}

func (p *PSql) Get(path string) gjson.Result {
	return gjson.Get(p.Content, path)
}

func (p *PSql) CheckFormat() bool {
	if !p.checkModel() {
		return false
	}

	return true
}

func (p *PSql) checkModel() bool {
	modelResult := p.Get("model")
	if err := json.Unmarshal([]byte(modelResult.String()), &p.pjson.model); err != nil {
		logger.Error("InitFile json.Unmarshal fail:%v", err)
		return false
	}

	return true
}

func GetModel() Model {
	return p.pjson.model
}

// GetFieldType 获取 model 字段类型
func GetFieldType(field string) string {
	res := ""
	if utils.FindStrIgnoreCaseInSlice(field, p.pjson.model.IntAttributes) {
		res = INT
	} else if utils.FindStrIgnoreCaseInSlice(field, p.pjson.model.StringAttributes) {
		res = STRING
	}

	if res == "" {
		logger.Error("GetFieldType ", field+" not found in model ", p.pjson.model)
	}

	return res
}

// GetSingleDeleteFiledJSON 获取删除操作的单一字段
func GetSingleDeleteFiledJSON() []string {
	gjsonRst := p.Get("Operate.D.single").Array()
	var rst = make([]string, 0)
	for _, r := range gjsonRst {
		rst = append(rst, r.String())
	}

	return rst
}

func GetMultipleDeleteFiledJSON() {

}

func GetFileName() string {
	return p.FileName
}
