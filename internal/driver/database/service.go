package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize a new connection with Database
func New(connection string) *gorm.DB {

	// Connect with database
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "gorm"}).Fatal(err)
	}

	return db
}
