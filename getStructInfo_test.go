package main

import "testing"

func TestGetStructureInfo(t *testing.T) {
	_, _ = GetStructureInfo(struct {
		Test string `bson:"test" gorm:"column:recode_state" desc:"测试"`
	}{})
}
