package main

import (
	"chamada-pagamento-system/internal/infra/repositories"
	"chamada-pagamento-system/tutorial"
	"context"
	"fmt"
)

func main() {
	// mux := http.NewServeMux()
	// httpserver.RegisterHandlers(mux)

	// migrations.PostgresMigrate()
	// migrations.DB.Migrator().DropTable(&entities.Associated{})
	// migrations.DB.AutoMigrate(&entities.Associated{})

	// log.Println("servidor inicializado em :8080")
	// if err := http.ListenAndServe(":8080", mux); err != nil {
	// 	log.Fatal("Erro ao iniciar servidor:", err)
	// }
	db := repositories.PgxConnect()
	queries := tutorial.New(db)
	ctx := context.Background()

	err := queries.CreateAssoc(ctx, tutorial.CreateAssocParams{
		Cpf:           123,
		Name:          "edu",
		DateBirth:     "1010",
		MaritalStatus: "lala",
	})
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	err = queries.GetAssoc(ctx)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
}
