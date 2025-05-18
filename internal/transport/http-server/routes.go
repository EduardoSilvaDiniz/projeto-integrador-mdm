package httpserver

import (
	"net/http"

	"chamada-pagamento-system/internal/transport/http-server/handlers"
)

func pingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", pingPong)
	mux.HandleFunc("GET /associated", handlers.GetAssociated)
	mux.HandleFunc("POST /associated", handlers.CreateAssociated)
	// mux.HandleFunc("PUT /associated", handlers.MapEndpointsToAssoc)
	mux.HandleFunc("DELETE /associated/{cpf}", handlers.DeleteAssoc)
}
