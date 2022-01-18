package main

import (
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/matrixcloud/proxy-pool/core"
	"github.com/matrixcloud/proxy-pool/server"
	"github.com/spf13/viper"
)

func main() {
	// create logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// Read configuration
	wd, _ := os.Getwd()
	cfg := path.Join(wd, "conf", "app.yaml")
	viper.SetConfigFile(cfg)
	if viper.ReadInConfig() != nil {
		log.Error().Msgf("Faild read config file from %s", cfg)
		return
	}
	// create pool
	pool := core.NewPool(
		viper.GetInt("pool.minThreshold"),
		viper.GetInt("pool.maxThreshold"),
	)

	// create scheduler
	shed := core.NewScheduler(pool)

	// Config api server
	port := viper.GetInt("server.port")
	svr := server.NewServer(pool, port)

	// start scheduler & server
	shed.Start()
	svr.Start()
}
