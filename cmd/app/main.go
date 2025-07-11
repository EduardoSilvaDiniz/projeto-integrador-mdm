package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/web"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func run() (*db.Queries, error) {
	ctx := context.Background()
	sqlite, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	if _, err := sqlite.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := db.New(sqlite)

	return queries, nil
}

func main() {
	log.Println("iniciando conexão com banco...")
	queries, err := run()
	if err != nil {
		log.Panic(err)
		return
	}
	log.Println("done")

	log.Println("iniciando servidor http...")
	mux := http.NewServeMux()
	web.CreateRouter(mux, queries)

	log.Println("servidor inicializado em :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
		return
	}
}
