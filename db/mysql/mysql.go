package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"template/config"
	"template/db/mysql/table"
)

/*
这个用于初始化数据库连接，使用gorm
*/

var Db *gorm.DB

func Init(cfg *config.MySQLConfig) (err error) {
	//1.连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//没连上就return
	if err != nil {
		fmt.Printf("mysql connect failed,err:%v\n", err)
		return
	}

	//2.设置连接池
	var sqlDB *sql.DB
	sqlDB, err = Db.DB()
	if err != nil {
		fmt.Printf(".DB() failed,err:%v\n", err)
		return
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) //最大连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns) //空闲连接数

	//3.自己去迁移表
	err = Db.AutoMigrate(&table.ExampleTable{})

	if err != nil {
		fmt.Println("迁移表失败")
		return
	}
	return
}

// Close 关闭数据库连接
func Close() {
	sqlDB, _ := Db.DB()
	err := sqlDB.Close()
	if err != nil {
		fmt.Printf(" close failed,err:%v\n", err)
		return
	}
}
