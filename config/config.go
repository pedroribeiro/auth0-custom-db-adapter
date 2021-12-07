package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	LOGRUS     = "logrus"
	SQLDB      = "postgres"
	TIMESERIES = "timestream"
	GIN        = "gin"
)

func ReadConfig(filename string) error {
	if os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load(filename)
		if err != nil {
			return errors.Wrap(err, "")
		}
	}

	return nil
}
