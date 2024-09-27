package config

import (
	"log"
	"os"
	"strconv"
)

func requiredStr(key string) (value string) {
	var ok bool
	if value, ok = os.LookupEnv(key); ok {
		return value
	} else {
		log.Panicf("Environment variable %v must be set.", key)
		return value
	}
}

func optionalStr(key string, defaultValue string) (value string) {
	var ok bool
	if value, ok = os.LookupEnv(key); ok {
		if !ok {
			return defaultValue
		}
		return value
	} else {
		return defaultValue
	}
}

func optionalUint(key string, defaultValue uint64) uint64 {
	if str, ok := os.LookupEnv(key); ok {
		if !ok {
			return defaultValue
		}
		value, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return defaultValue
		}
		return value
	} else {
		return defaultValue
	}
}
