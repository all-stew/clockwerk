package mysql

import (
	"clockwerk/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	conf := config.Config.MySQLConfig
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	engine, _ = xorm.NewEngine("mysql", dsn)
	err := engine.Ping()
	if err != nil {
		fmt.Println("mysql连接数据库失败")
	} else {
		fmt.Println("mysql连接数据库成功!!!")
		engine.SetMaxOpenConns(conf.MaxOpenConns)
		engine.SetMaxIdleConns(conf.MaxIdleConns)
	}
}
func GetDb() *xorm.Engine {
	return engine
}

// Close 关闭MySQL连接
func Close() {
	_ = engine.Close()
}
