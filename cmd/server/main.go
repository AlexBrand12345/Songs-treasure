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
	fmt.Println("Я запустился")

	err = config.LoadConfig()
	if err != nil {
		log.Fatalf("godotenv failed to load env file: %v", err)
	}

	logging.StartLogrus(config.LOG_LEVEL)

	DB, err := db.DbConnect(config.DB_HOST,
		config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_NAME)
	if err != nil {
		logging.Default.Fatalf("Couldn`t connect to DB")
	}

	service := service.NewService(DB)
	controller := controller.NewController(service)

	logging.Default.Infof("Starting server. Port:%s", config.PORT)
	logging.Default.Error(http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), router.Router(controller)).Error())
}
