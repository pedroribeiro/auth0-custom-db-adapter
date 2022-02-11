package migrations

import (
	"github.com/pedroribeiro/users/internal/domain/entity"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(
		&entity.User{},
	)
}
