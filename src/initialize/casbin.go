package initialize

import (
	"clockwerk/src/global"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

/*
   说明：初始化 Casbin
*/
func MysqlCasbin() {
	// 初始化数据库适配器，并自定义表名，默认 casbin_rule
	a, err := gormadapter.NewAdapterByDBUseTableName(global.DB, global.Conf.Mysql.TablePrefix, "casbin_rule")
	if err != nil {
		panic(fmt.Sprintf("Casbin数据库连接失败：%s", err.Error()))
	}

	// 读取配置文件，需要避免多次读取
	config, err := global.ConfBox.Find(global.Conf.Casbin.Model)
	cabinModel := model.NewModel()

	// 从字符串中加载配置
	err = cabinModel.LoadModelFromText(string(config))
	if err != nil {
		panic(fmt.Sprintf("Casbin配置读取失败：%s", err.Error()))
	}

	e, err := casbin.NewEnforcer(cabinModel, a)
	if err != nil {
		panic(fmt.Sprintf("Casbin配置加载失败：%s", err.Error()))
	}

	// 加载策略
	err = e.LoadPolicy()
	if err != nil {
		panic(fmt.Sprintf("Casbin策略加载失败：%s", err.Error()))
	}

	// 全局设置
	//global.CasbinEnforcer = e
	log.Println("Casbin初始化完成")
}
