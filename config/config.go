package config

type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Mode        string      `mapstructure:"mode"`
	Port        int         `mapstructure:"port"`
	Version     string      `mapstructure:"version"`
	JWTConfig   JWTConfig   `mapstructure:"jwt"`
	RedisConfig RedisConfig `mapstructure:"redis"`
	LogConfig   LogConfig   `mapstructure:"logger"`
	MySQLConfig MySQLConfig `mapstructure:"mysql"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password_util"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type JWTConfig struct {
	Key string `mapstructure:"key" json:"key"`
}
