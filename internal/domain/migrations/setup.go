package migrations

import (
	"github.com/elvenworks/users/internal/domain/entity"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(
		&entity.User{},
	)
}
