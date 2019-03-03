package main

import (
	"log"

	"monita/config"
	"monita/handlers"
	"monita/storage"

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
