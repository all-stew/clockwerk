package common

import "go.uber.org/zap/zapcore"

// 配置总入口
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
	Logs   LogsConfiguration   `mapstructure:"logs" json:"logs"`
	Mysql  MysqlConfiguration  `mapstructure:"mysql" json:"mysql"`
	Jwt    JwtConfiguration    `mapstructure:"jwt" json:"jwt"`
	Casbin CasbinConfig        `mapstructure:"casbin" json:"casbin"`
	Redis  RedisConfiguration  `mapstructure:"redis" json:"redis"`
}

// 系统配置项
type SystemConfiguration struct {
	Name       string `mapstructure:"name" json:"name"`
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	Mode       string `mapstructure:"mode" json:"mode"`
	ApiPrefix  string `mapstructure:"api-prefix" json:"apiPrefix"`
	ApiVersion string `mapstructure:"api-version" json:"apiVersion"`
	Avatar     string `mapstructure:"avatar" json:"avatar"`
}

// 日志配置项
type LogsConfiguration struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Colorful   bool          `mapstructure:"colorful" json:"colorful"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

// 数据库配置项
type MysqlConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     string `mapstructure:"database" json:"database"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Charset      string `mapstructure:"charset" json:"charset"`
	Collation    string `mapstructure:"collation" json:"collation"`
	Query        string `mapstructure:"query" json:"query"`
	TablePrefix  string `mapstructure:"table-prefix" json:"tablePrefix"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"maxIdleTime"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
	LogLevel     int    `mapstructure:"log-level" json:"logLevel"`
}

// Jwt 配置项目
type JwtConfiguration struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

// Casbin 配置
type CasbinConfig struct {
	Model string `mapstructure:"model" json:"model"`
}

type RedisConfiguration struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}
