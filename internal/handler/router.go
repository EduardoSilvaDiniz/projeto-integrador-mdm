package handler

import (
	"chamada-pagamento-system/internal/database"
	"chamada-pagamento-system/internal/service"
	"net/http"
)

func PingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateRouter(mux *http.ServeMux, queries *database.Queries) {
	associatedService := service.NewAssociatedService(queries)
	associatedHandler := NewAssociatedHandler(associatedService)

	mux.HandleFunc("GET /ping", PingPong)

	mux.HandleFunc("GET /associated", associatedHandler.List())
	mux.HandleFunc("POST /associated", associatedHandler.Create())
	mux.HandleFunc("DELETE /associated/{number_card}", associatedHandler.Delete())
}
