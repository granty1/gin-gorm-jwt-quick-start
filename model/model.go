package model

import (
	"fmt"
	"github.com/gin-cli/config"
	"github.com/gin-cli/init/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Packages used
//gorm 			 https://github.com/jinzhu/gorm
//connect driver https://github.com/jinzhu/gorm/dialects/mysql

//Connect used to manage mysql connect
type Connect struct {
	mysqlConnect *gorm.DB
}

var connect Connect

//Init to init mysql connect
func Init(dbConfig config.DBConfig) {
	var err error
	connect.mysqlConnect, err = gorm.Open(dbConfig.Type,
		fmt.Sprintf("%s:%s@/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.DBname, dbConfig.Config))
	if  err != nil {
		log.Log.Error("connect db fail, error,", err)
		panic("connect to db fail, for more see log")
	}
	connect.mysqlConnect.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	connect.mysqlConnect.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
}

//GetMysqlConn return a instance of mysql connect.
func GetMysqlConn() *gorm.DB {
	return connect.mysqlConnect
}
