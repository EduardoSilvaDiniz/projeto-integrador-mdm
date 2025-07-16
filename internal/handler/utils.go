package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func writeError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{"error": msg})
}

func writeOk(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func logInicial(r *http.Request) {
	ip := r.RemoteAddr
	ua := r.UserAgent()
	method := r.Method
	path := r.URL.Path

	slog.Debug("Requisição recebida",
		"ip", ip,
		"user_agent", ua,
		"method", method,
		"path", path,
	)
}

func serviceError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error(
		"Erro inesperado durante execução da operação",
		"path",
		r.URL.Path,
		"err",
		err,
	)
	writeError(
		w,
		"Ocorreu um erro ao processar sua solicitação. Tente novamente mais tarde.",
		http.StatusInternalServerError,
	)
}
