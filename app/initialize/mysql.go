package initialize

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	"clockwerk/app/repository/store"
	"clockwerk/pkg/dbutils"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
   说明：MySQL 数据库连接初始化
*/
func Mysql() {
	// 用于打印的连接
	dsnLog := fmt.Sprintf("%s:******@tcp(%s:%d)/%s?%s&charset=%s&collation=%s",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Database,
		global.Conf.Mysql.Query,
		global.Conf.Mysql.Charset,
		global.Conf.Mysql.Collation,
	)

	// 实际用于连接数据库的连接串
	dsn := strings.Replace(dsnLog, "******", global.Conf.Mysql.Password, 1)

	// 打印数据库连接串
	global.Log.Debug("打开连接（MySQL）：", dsnLog)

	// 打开数据库链接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // 连接信息
	}), &gorm.Config{
		QueryFields: true, // 解决查询全部字段可能不走索引的问题
	})

	// 错误退出
	if err != nil {
		message := fmt.Sprintf("数据库连接异常：%s", err.Error())
		global.Log.Error(message)
		panic(message)
	}

	// db
	global.DB = db
	// 获取数据库连接
	global.UserStore = store.NewUserStore(dbutils.NewConnection(db))

	log.Println("数据库初始化完成")

	// 数据同步
	TableAutoMigrate()
	InitData()
}

// 数据同步
func TableAutoMigrate() {
	err := global.DB.AutoMigrate(
		new(models.SysUser),             // 用户 数据表
		new(models.SysUserRoleRelation), // 用户-角色 数据表
		new(models.SysRole),             // 角色 数据表
		new(models.SysRoleMenuRelation), // 角色-菜单 数据表
		new(models.SysMenu),             // 菜单 数据表
		new(models.SysPreset),           // 策略集 数据表
		new(models.SysRolePreset),       // 角色-策略集 数据表
	)
	if err != nil {
		panic(fmt.Sprintf("数据库同步异常：%s", err.Error()))
	}
}
