package config

import (
	"fmt"
	"os"
)

func GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

func GetDbUri() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
}

func GetPort() string {
	return os.Getenv("LISTENING_PORT")
}
