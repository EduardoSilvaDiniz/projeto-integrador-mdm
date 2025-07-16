package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/web"

	_ "modernc.org/sqlite"
)

var (
	//go:embed schema.sql
	ddl string

	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
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

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("Panic recuperado", "error", rec)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode("erro interno do servidor")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := ":8080"
	slog.SetDefault(logger)
	slog.Info("Iniciado conexão com banco de dados sqlite")
	queries, err := run()
	if err != nil {
		slog.Error("falha na conexão com banco de dados", "err", err)
		return
	}

	slog.Info("Iniciando Servidor HTTP", "port", port)
	mux := http.NewServeMux()
	web.CreateRouter(mux, queries)

	if err := http.ListenAndServe(port, RecoverMiddleware(mux)); err != nil {
		slog.Error("Erro ao iniciar servidor:")
		return
	}
}
