package migrations

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	config = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
)

func PostgresMigrate() {
	database, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar ao banco de dados", err)
	}
	DB = database
}
