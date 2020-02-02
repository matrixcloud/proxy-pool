package main

import (
	"fmt"
	"github.com/matrixcloud/proxy-pool/db"
	"github.com/matrixcloud/proxy-pool/pool"
	"github.com/matrixcloud/proxy-pool/server"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
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

	// Config db
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetUint("db.port")
	dbPass := viper.GetString("db.pass")
	dbOptions := db.Options{
		Host: dbHost,
		Port: dbPort,
		Pass: dbPass,
	}
	conn := db.NewClient(&dbOptions)

	if !db.Test(conn) {
		log.Fatalf("Failed to connect to redis server: %s", fmt.Sprintf("%s:%d", dbHost, dbPort))
		return
	}

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
