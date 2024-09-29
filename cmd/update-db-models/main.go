package main

import (
	"log"
	"songs-treasure/internal/config"
	"songs-treasure/pkg/db"
	"songs-treasure/pkg/logging"

	"gorm.io/gen"
)

func main() {
	var err error

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

	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/db/tables_functions",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(DB.DB)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()

	logging.Default.Info("Generated tables from DB")
}
