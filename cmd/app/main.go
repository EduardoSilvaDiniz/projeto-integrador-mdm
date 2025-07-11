package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"log/slog"
	"net/http"
	"os"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/web"

	_ "modernc.org/sqlite"
)

var (
	//go:embed schema.sql
	ddl    string
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // ou slog.LevelDebug
	}))
)

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
	slog.SetDefault(logger)
	slog.Info("Iniciado conexão com banco de dados sqlite")
	queries, err := run()
	if err != nil {
		slog.Error("falha na conexão com banco de dados", "err", err)
		return
	}
	slog.Info("feito")

	slog.Info("Iniciando Servidor HTTP")
	mux := http.NewServeMux()
	web.CreateRouter(mux, queries)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Erro ao iniciar servidor:")
		return
	}
	slog.Info("servidor inicializado em :8080")
}
