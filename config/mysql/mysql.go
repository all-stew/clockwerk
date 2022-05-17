package mysql

import (
	"clockwerk/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var engine *gorm.DB

func Init(conf *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("mysql连接数据库失败 %s", err)
		return
	}
	return
}

func GetDb() *gorm.DB {
	return engine
}
