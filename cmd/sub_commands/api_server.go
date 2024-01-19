package sub_commands

import (
	"fmt"
	"github.com/jinglanghe/go-start/internal/app"
	"github.com/jinglanghe/go-start/utils/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "0.0.1"
	appId   = "go-start"
)

// ApiServer Api_server represents the alert command
var ApiServer = &cobra.Command{
	Use:   "api_server",
	Short: "api_server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("go-start project is starting")
		newApp, cleanFunc, err := app.Build()
		if err != nil {
			log.Fatal().Err(err).Msg("app build failed")
		}

		log.Info().Interface("newApp", newApp).Interface("cleanFunc", cleanFunc).Msg("init app successfully")

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			log.Info().Str(appId, fmt.Sprintf("get a signal %s", s.String())).Send()
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				cleanFunc()
				log.Info().Str(appId, fmt.Sprintf("[version: %s] exit", version)).Send()
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	},
}
