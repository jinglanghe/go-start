package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinglanghe/go-start/internal/cache"
	"github.com/jinglanghe/go-start/internal/config"
	"gorm.io/gorm"
)

var ApplicationSet = wire.NewSet(wire.Struct(new(Application), "*"))

type Application struct {
	Config    *config.AppConfig
	Db        *gorm.DB
	Cache     cache.Adapter
	GinEngine *gin.Engine
}
