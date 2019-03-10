package storage

import (
	"monita/storage/observable"
	"monita/storage/user"

	"github.com/jinzhu/gorm"
)

// Init runs initialization code for all storage
func Init(db *gorm.DB) {
	user.Init(db)
	observable.Init(db)

	go user.Worker()
}
