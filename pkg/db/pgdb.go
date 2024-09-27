package db

import (
	"fmt"
	"songs-treasure/pkg/logging"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgres struct {
	*gorm.DB
}

func DbConnect(dbhost, dbport, dbuser, dbpass, dbname string) (pgdb *postgres, err error) {
	var connData string = fmt.Sprintf(
		"host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='disable'",
		dbhost, dbport, dbuser, dbpass, dbname,
	)
	logging.Default.Debugf(connData)
	db, err := gorm.Open("postgres", connData)
	if err != nil {
		fmt.Println("error")
		logging.Default.Errorf("DB connection error: %v", err)
		return
	}

	pgdb = &postgres{db}
	logging.Default.Info("Connected to DB")

	return
}

func DbMigrate(DB *postgres) {
	DB.AutoMigrate()

	logging.Default.Info("AutoMigration completed")
}
