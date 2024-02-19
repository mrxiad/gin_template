package table

/*
这个用于定义数据库的表
*/

import (
	"gorm.io/gorm"
)

type ExampleTable struct {
	gorm.Model
	// 这里定义表的字段
}

// TableName 这里定义表的名字
func (ExampleTable) TableName() string {
	return "example_service"
}
