package psql

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"path/filepath"
	"psql/pkg/logger"
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

	//pjsonByte, err := json.Marshal(fByte)
	//// Check json format
	//if err != nil {
	//	logger.Fatal("InitFile json.Marshal fail:%v", err)
	//}

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

func GetFileName() string {
	return p.FileName
}
