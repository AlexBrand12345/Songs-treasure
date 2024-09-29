package db

import (
	"fmt"
	"songs-treasure/pkg/db/model"
	"songs-treasure/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pg struct {
	*gorm.DB
}

type DBParams struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func DbConnect(dbparams DBParams) (pgdb *pg, err error) {
	var connData string = fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		dbparams.DBHost, dbparams.DBUser, dbparams.DBPass, dbparams.DBName, dbparams.DBPort,
	)
	logging.Default.Debugf(connData)

	db, err := gorm.Open(postgres.Open(connData), &gorm.Config{})
	if err != nil {
		logging.Default.Errorf("DB connection error: %v", err)
		return
	}

	pgdb = &pg{db}
	logging.Default.Info("Connected to DB and generated missing structs")

	return
}

func DbMigrate(db *pg) (err error) {
	migration := db.Migrator()

	logging.Default.Debugf("db: has groups - %v, has songs - %v, has songs_verses - %v",
		migration.HasTable("groups"), migration.HasTable("songs"), migration.HasTable("songs_verses"))

	if migration.HasTable("groups") &&
		migration.HasTable("songs") &&
		migration.HasTable("songs_verses") {
		logging.Default.Infof("No need to migrate")
		return
	}

	err = db.AutoMigrate(
		&model.Group{},
		&model.Song{},
		&model.SongsVerse{},
	)
	if err != nil {
		logging.Default.Errorf("DB migration error: %v", err)
		return
	}

	logging.Default.Info("AutoMigration completed")

	return
}
