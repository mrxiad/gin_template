package impl

import "template/dao"

/*
实现dao层的接口,操作数据库
*/
import (
	"gorm.io/gorm"
	"template/db/mysql/table"
)

type ExampleDaoImpl struct {
	db *gorm.DB
}

// NewExampleDaoImpl 这个用于依赖注入，但是懒得写了
func NewExampleDaoImpl(db *gorm.DB) dao.ExampleDao {
	return &ExampleDaoImpl{db: db}
}

// QueryExampleById 这里面操作数据库，写sql
func (o *ExampleDaoImpl) QueryExampleById(id uint) (table.ExampleTable, error) {
	var ExampleTable table.ExampleTable
	result := o.db.First(&ExampleTable, id)
	if result.Error != nil {
		return table.ExampleTable{}, result.Error
	}
	return ExampleTable, nil
}
