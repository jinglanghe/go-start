//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/jinglanghe/go-start/internal/config"
)

func Build() (*Application, func(), error) {
	wire.Build(
		config.Init,
		ApplicationSet,
	)
	return new(Application), nil, nil
}
