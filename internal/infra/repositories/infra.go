package repositories

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"


func PostgresMigrate() *gorm.DB {
	database, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar ao banco de dados", err)
	}
	return database
}
