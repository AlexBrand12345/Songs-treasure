package main

import (
	"fmt"
	"log"
	"net/http"
	"songs-treasure/controller"
	"songs-treasure/internal/config"
	"songs-treasure/internal/service"
	"songs-treasure/pkg/db"
	"songs-treasure/pkg/logging"
	"songs-treasure/router"
)

func main() {
	var err error
	fmt.Println("Hello there!")

	err = config.LoadConfig()
	if err != nil {
		log.Fatalf("godotenv failed to load env file: %v", err)
	}

	logging.StartLogrus(config.LOG_LEVEL)

	DB, err := db.DbConnect(db.DBParams{
		DBHost: config.DB_HOST,
		DBPort: config.DB_PORT,
		DBUser: config.DB_USER,
		DBPass: config.DB_PASS,
		DBName: config.DB_NAME,
	})
	if err != nil {
		logging.Default.Fatalf("Couldn`t connect to DB")
	}

	err = db.DbMigrate(DB)

	if err != nil {
		logging.Default.Fatalf("Couldn`t complete DB migration")
	}

	service := service.NewService(DB)
	controller := controller.NewController(service)

	logging.Default.Infof("Starting server. Port:%s", config.PORT)
	logging.Default.Error(http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), router.Router(controller)).Error())
}
