package main

import (
	"reflect"
)

type DatabaseInfo struct {
	Name string
	Type string
	Bson string
	Gorm string
	Desc string
}

func GetStructureInfo(p interface{}) ([]*DatabaseInfo, error) {
	pt := reflect.TypeOf(p)
	infos := make([]*DatabaseInfo, 0)
	if pt.Kind() == reflect.Struct {
		fieldNum := pt.NumField()
		for i := 0; i < fieldNum; i++ {
			infos = append(infos, &DatabaseInfo{
				Name: pt.Field(i).Name,            //字段名
				Type: pt.Field(i).Type.Name(),     //类型
				Bson: pt.Field(i).Tag.Get("bson"), //标记内容
				Gorm: pt.Field(i).Tag.Get("gorm"),
				Desc: pt.Field(i).Tag.Get("desc"),
			})
		}
	}

	return infos, nil
}
