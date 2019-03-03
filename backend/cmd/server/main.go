package main

import (
	"log"

	"monita-backend/config"
	"monita-backend/handlers"
	"monita-backend/storage"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", config.DBConnectionString)

	if err != nil {
		log.Fatalln(err)
	}

	storage.Init(db)
	handlers.Listen()
}
