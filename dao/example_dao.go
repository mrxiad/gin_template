package dao

/*
用于定义操作数据库的接口
*/
import (
	"template/db/mysql/table"
)

type ExampleDao interface {

	// QueryExampleById 参数是id，返回值是表+错误
	QueryExampleById(id uint) (table.ExampleTable, error)
}
