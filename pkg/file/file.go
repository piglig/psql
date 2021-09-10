package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"psql/pkg/components"
	"psql/pkg/logger"
	"psql/pkg/psql"
	"psql/pkg/utils"
	"strings"
)

type PFile struct {
	TargetPath     string
	TargetFileName string
	Content        string
}

var pfile = new(PFile)

func InitOutFileConfig(targetPath, targetFileName string) {
	pfile.TargetPath = targetPath
	pfile.TargetFileName = targetFileName
}

func (p *PFile) generatePackage(pkg string) {
	p.Content += fmt.Sprintf("package %s\n\n", pkg)
}

func (p *PFile) generateModel() {
	filename := psql.GetFileName()
	filePair := strings.Split(filename, ".")
	modelName := utils.FirstLetterToUpper(filePair[0])

	model := psql.GetModel()
	// model definition begin
	p.Content += fmt.Sprintf("type %s struct {\n", modelName)

	for _, attribute := range model.IntAttributes {
		p.Content += fmt.Sprintf("\t%s int\n", utils.FirstLetterToUpper(attribute))
	}

	for _, attribute := range model.StringAttributes {
		p.Content += fmt.Sprintf("\t%s string\n", utils.FirstLetterToUpper(attribute))
	}

	// model definition end
	p.Content += fmt.Sprintf("}\n\n")

	p.generateModelTable(modelName)
}

func (p *PFile) generateDao() {
	filename := psql.GetFileName()
	filePair := strings.Split(filename, ".")
	daoName := utils.FirstLetterToUpper(filePair[0])

	// dao definition begin
	p.Content += fmt.Sprintf("type %sDao struct {}\n\n", daoName)
	p.Content += components.GenerateDaoObjFun(daoName)

	deleteSingleColumns := psql.GetSingleDeleteFiledJSON()
	for _, column := range deleteSingleColumns {
		p.Content += components.GenerateDaoDeleteByKeyFunc(daoName, column, psql.GetFieldType(column))
	}

	deleteMultipleColumns := psql.GetMultipleDeleteFiledJSON()
	for _, columns := range deleteMultipleColumns {
		p.Content += components.GenerateDaoDeleteByKeysFunc(daoName, columns)
	}

	p.Content += components.GenerateDaoCreateFunc(daoName)

	//p.generateModelTable(modelName)
}

func (p *PFile) generateModelTable(modelName string) {
	p.Content += fmt.Sprintf("func (*%s) TableName() string {\n", modelName)
	p.Content += fmt.Sprintf("\treturn \"%s\"\n", strings.ToLower(modelName))
	p.Content += fmt.Sprintf("}\n\n")
}

func (p *PFile) createModelFile() bool {
	pfile.generatePackage("model")
	pfile.generateModel()

	filePathPrefix := "/model/"
	pfile.createFile(filePathPrefix)
	return true
}

func (p *PFile) createDaoFile() bool {

	pfile.generatePackage("dao")
	pfile.generateDao()

	filePathPrefix := "/dao/"
	pfile.createFile(filePathPrefix)

	return true
}

func (p *PFile) createFile(targetPathPrefix string) bool {
	// if path not exist, mkdir it
	fp := pfile.TargetPath + targetPathPrefix
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		err := os.Mkdir(fp, 777)
		if err != nil {
			logger.Error("CreateFile mkdir err:%v", err)
			return false
		}
	}

	err := ioutil.WriteFile(fp+pfile.TargetFileName, []byte(pfile.Content), 755)
	if err != nil {
		logger.Error("CreateFile err:%v", err)
		return false
	}
	pfile.Content = ""

	return true
}

func Run() {
	pfile.createModelFile()
	pfile.createDaoFile()
}
