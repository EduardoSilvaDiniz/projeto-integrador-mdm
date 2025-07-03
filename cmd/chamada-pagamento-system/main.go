package main

import (
	"chamada-pagamento-system/internal/endpoint"
	"log"
	"net/http"
)

func main() {
	// TODO Fazer o migration funciona
	// deployments.Migration(conn)
	mux := http.NewServeMux()
	endpoint.CreateEndpoints(mux)

	log.Println("servidor inicializado em :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
