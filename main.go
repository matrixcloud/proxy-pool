package main

import (
	"log"
	"os"
	"path"

	"github.com/matrixcloud/proxy-pool/db"
	"github.com/matrixcloud/proxy-pool/pool"
	"github.com/matrixcloud/proxy-pool/server"
	"github.com/spf13/viper"
)

func main() {
	// Read configuration
	wd, _ := os.Getwd()
	cfg := path.Join(wd, "conf", "app.yaml")
	viper.SetConfigFile(cfg)
	if viper.ReadInConfig() != nil {
		log.Fatalf("Faild read config file from %s", cfg)
		return
	}
	conn := db.NewClient()

	// Config proxy pool
	poolOptions := pool.Options{
		MaxThreshold:     viper.GetInt("pool.maxThreshold"),
		MinThreshold:     viper.GetInt("pool.minThreshold"),
		CheckInterval:    viper.GetInt("pool.checkInterval"),
		ValidateInterval: viper.GetInt("pool.validateInterval"),
	}
	p := pool.NewPool(conn, poolOptions)
	p.Start()

	// Config api server
	port := viper.GetInt("server.port")
	s := server.NewServer(conn, port)
	s.Start()
}
