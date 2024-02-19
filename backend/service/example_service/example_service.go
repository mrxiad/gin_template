package example_service

import (
	"template/dao"
	"template/db/mysql/table"
)

type ExampleService struct {
	ExampleDao dao.ExampleDao //这里用接口而不是用具体的实现类
}

// QueryExampleById 返回的结果尽量是表结构
func (o *ExampleService) QueryExampleById(id uint) (table.ExampleTable, error) {
	return o.ExampleDao.QueryExampleById(id)
}
