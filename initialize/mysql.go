package initialize

import (
	"clockwerk/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() {
	mysqlConfig := global.ServerSetting.MySQLConfig
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DB)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db
}
