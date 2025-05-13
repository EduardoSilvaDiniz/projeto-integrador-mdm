package main

import (
	"log"
	"net/http"

	"chamada-pagamento-system/db"
	"chamada-pagamento-system/internal/domain"
	"chamada-pagamento-system/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.RegisterHandlers(mux)

	db.Connect()
	db.DB.AutoMigrate(&domain.Associated{})

	log.Println("servidor inicializado em :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
