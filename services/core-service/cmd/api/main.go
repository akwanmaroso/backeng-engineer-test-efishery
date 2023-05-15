package main

import (
	"log"
	"os"

	"github.com/akwanmaroso/backend-efishery-test/core-service/config"
	"github.com/akwanmaroso/backend-efishery-test/core-service/internal/api"
	"github.com/akwanmaroso/backend-efishery-test/core-service/utils"
)

func main() {
	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	srv := api.NewApi(cfg)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
