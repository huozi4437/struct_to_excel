package main

import (
	"fmt"
	"getDatabaseDesc/object"
	"github.com/tealeg/xlsx"
	"log"
)

var AllCollectMap map[string]interface{}

func main() {
	AllCollectMap = make(map[string]interface{}, 0)
	AllCollectMap["student"] = object.StudentInfo{}
	AllCollectMap["test_scores"] = object.TestScores{}

	CreateXlsx()
}

func CreateXlsx() {
	var file *xlsx.File
	var err error

	file = xlsx.NewFile()
	defer func() {
		err = file.Save("文件名称.xlsx") //保存，输出表格文件
		if err != nil {
			log.Println("file save err:", err)
		}
	}()

	for key, collec := range AllCollectMap {
		sheet, err := file.AddSheet(key) //添加表单，key：表单名
		if err != nil {
			log.Printf("AddSheet %s err:%v", key, err)
			continue
		}

		collectInfo, err := GetStructureInfo(collec)
		if err != nil {
			log.Printf("GetStructureInfo %s err:%v", key, err)
			continue
		}
		row := sheet.AddRow()
		row.AddCell().SetString(fmt.Sprintf("表名：%s", key))
		row = sheet.AddRow()
		row.AddCell().SetString("变量名")
		row.AddCell().SetString("类型")
		row.AddCell().SetString("MongoDB字段")
		row.AddCell().SetString("SQLite字段")
		row.AddCell().SetString("备注")
		for _, info := range collectInfo {
			row := sheet.AddRow()
			row.AddCell().SetString(info.Name)
			row.AddCell().SetString(info.Type)
			row.AddCell().SetString(info.Bson)
			row.AddCell().SetString(info.Gorm)
			row.AddCell().SetString(info.Desc)
		}
		//设置宽度
		_ = sheet.SetColWidth(0, 0, 25)
		_ = sheet.SetColWidth(1, 1, 10)
		_ = sheet.SetColWidth(2, 100, 25)
	}

}
