//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/jinglanghe/go-start/internal/config"
	"github.com/jinglanghe/go-start/internal/database"
)

func Build() (*Application, func(), error) {
	wire.Build(
		config.Init,
		database.InitDb,
		ApplicationSet,
	)
	return new(Application), nil, nil
}
