package config

import (
	"github.com/joho/godotenv"
)

var (
	LOG_LEVEL uint8
	PORT      string

	MUSIC_INFO_URL string

	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_PORT string
	DB_HOST string
)

func LoadConfig() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}

	LOG_LEVEL = uint8(optionalUint("LOG_LEVEL", 3))
	PORT = requiredStr("PORT")

	MUSIC_INFO_URL = requiredStr("MUSIC_INFO_URL")

	DB_USER = requiredStr("DB_USER")
	DB_PASS = requiredStr("DB_PASS")
	DB_NAME = requiredStr("DB_NAME")
	DB_PORT = requiredStr("DB_PORT")
	DB_HOST = requiredStr("DB_HOST")

	return
}
