package main

import (
	"log"
	"net/http"


	httpserver "chamada-pagamento-system/internal/transport/http-server"
)

func main() {
	mux := http.NewServeMux()
	httpserver.RegisterHandlers(mux)

	// migrations.PostgresMigrate()
	// migrations.DB.Migrator().DropTable(&entities.Associated{})
	// migrations.DB.AutoMigrate(&entities.Associated{})

	log.Println("servidor inicializado em :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
