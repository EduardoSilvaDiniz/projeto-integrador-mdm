package main

import (
	"chamada-pagamento-system/deployments"
	"chamada-pagamento-system/internal/infra/repositories"
	httpserver "chamada-pagamento-system/internal/transport/http-server"
	"log"
	"net/http"
)

func main() {
	conn := repositories.PgxConnect()
	deployments.Migration(conn)
	// queries := db.New(conn)
	// ctx := context.Background()
	//
	// if err := queries.CreateAssoc(ctx, db.CreateAssocParams{
	// 	Cpf:           123,
	// 	Name:          "edu",
	// 	DateBirth:     "1010",
	// 	MaritalStatus: "lala",
	// }); err != nil {
	// 	fmt.Println("ERROR: ", err)
	// 	return
	// }
	//
	// err := queries.GetAssoc(ctx)
	// if err != nil {
	// 	fmt.Println("ERROR: ", err)
	// 	return
	// }

	mux := http.NewServeMux()
	httpserver.RegisterHandlers(mux)

	log.Println("servidor inicializado em :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
