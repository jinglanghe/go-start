package main

import (
	"github.com/jinglanghe/go-start/internal/app"
	"github.com/jinglanghe/go-start/utils/log"
)

const (
	ver   = "0.0.1"
	appid = "go-start"
)

func main() {
	log.Info().Msg("go-start project is starting")
	newApp, cleanFunc, err := app.Build()
	if err != nil {
		log.Fatal().Err(err).Msg("app build failed")
	}

	log.Info().Interface("newApp", newApp).Interface("cleanFunc", cleanFunc).Msg("init app successfully")
}
