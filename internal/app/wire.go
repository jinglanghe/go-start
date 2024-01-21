//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/jinglanghe/go-start/internal/cache"
	"github.com/jinglanghe/go-start/internal/config"
	"github.com/jinglanghe/go-start/internal/database"
	"github.com/jinglanghe/go-start/internal/web"
)

func Build() (*Application, func(), error) {
	wire.Build(
		config.Init,
		database.InitDb,
		cache.Init,
		web.InitGinEngine,
		ApplicationSet,
	)
	return new(Application), nil, nil
}
