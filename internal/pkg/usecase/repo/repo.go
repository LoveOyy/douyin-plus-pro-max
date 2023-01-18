package repo

import (
	"fmt"

	"douyin/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	config := config.Data()
	var err error

	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4", config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Addr,
		config.Mysql.DBName,
	)))
	if err != nil {
		panic(err)
	}

}
