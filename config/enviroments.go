package config

import (
	"fmt"
	"os"
)

func GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

func GetDbDsn() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
}
