package main

import (
	"nongki/config"
	"nongki/internal/router"
	"nongki/pkg/db"
	"nongki/pkg/log"
)

func main() {
	log.InitLogger()
	logger := log.GetLogger()

	cfg := config.LoadConfig()
	db, err := db.ConnectDB(cfg)
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	// App Init
	appConfig := config.AppConfig{
		Db:     db,
		Logger: logger,
	}

	router.Router(appConfig)
}
