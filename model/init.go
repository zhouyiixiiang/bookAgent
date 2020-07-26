package model

import (
	"config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() (err error) {
	DB, err = gorm.Open("mysql", config.Config.MysqlSetting[config.SrvName].MysqlConn)
	if err != nil {
		fmt.Println("gorm.Open( mysql err: ", err)
	}
	DB.DB().SetMaxOpenConns(config.Config.MysqlSetting[config.SrvName].MysqlConnectPoolSize)
	DB.DB().SetMaxIdleConns(config.Config.MysqlSetting[config.SrvName].MysqlConnectPoolSize / 2)
	//err =DB.DB().Ping()
	migration()
	return
}
