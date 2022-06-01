package global

import (
	"clockwerk/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ServerSetting config.ServerConfig
	Lg            *zap.Logger
	Trans         ut.Translator
	DB            *gorm.DB
	Redis         *redis.Client
)
