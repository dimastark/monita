package storage

import (
	"monita-backend/storage/observable"
	"monita-backend/storage/user"

	"github.com/jinzhu/gorm"
)

// Init runs initialization code for all storage
func Init(db *gorm.DB) {
	user.Init(db)
	observable.Init(db)
}
