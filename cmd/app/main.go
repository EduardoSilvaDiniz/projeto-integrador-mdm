package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/web"

	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func run() (*database.Queries, error) {
	ctx := context.Background()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := database.New(db)

	return queries, nil
}

func main() {
	log.Println("iniciando conexão com banco...")
	queries, err := run()
	if err != nil {
		log.Panic(err)
	}
	log.Println("done")

	log.Println("iniciando servidor http...")
	mux := http.NewServeMux()
	web.CreateRouter(mux, queries)

	log.Println("servidor inicializado em :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}

	log.Println("done")
}
