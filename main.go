package main

import (
	"chamada-pagamento-system/internal/domain"
	"chamada-pagamento-system/internal/handler"
	"chamada-pagamento-system/internal/infra"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler.RegisterHandlers(mux)
	db := infra.StartDb(10)
	p := domain.Associated{Cpf: "10210"}
	db[1] = p
	fmt.Println(db[1])

	log.Println("servidor inicializado em :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
