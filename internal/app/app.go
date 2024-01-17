package app

import (
	"github.com/google/wire"
	"github.com/jinglanghe/go-start/internal/config"
)

var ApplicationSet = wire.NewSet(wire.Struct(new(Application), "*"))

type Application struct {
	Config *config.AppConfig
	// TODO
	//Engine *gin.Engine
	//Db     *gorm.DB
}
