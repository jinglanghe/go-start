package app

import (
	"github.com/google/wire"
	"github.com/jinglanghe/go-start/internal/config"
	"gorm.io/gorm"
)

var ApplicationSet = wire.NewSet(wire.Struct(new(Application), "*"))

type Application struct {
	Config *config.AppConfig
	Db     *gorm.DB
	// TODO
	//Engine *gin.Engine

}
