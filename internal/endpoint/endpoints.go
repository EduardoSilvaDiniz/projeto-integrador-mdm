package endpoint

import (
	"chamada-pagamento-system/internal/database"
	"chamada-pagamento-system/internal/endpoint/handler"
	"net/http"
)

func pingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

type Handlers struct {
	AssociatedController handler.AssociatedController
}

func NewHandlers(associatedService handler.AssociatedController) *Handlers {
	return &Handlers{AssociatedController: associatedService}
}

func CreateEndpoints(mux *http.ServeMux, queries *database.Queries) {
	associatedService := handler.NewAssociatedService(queries)
	handlers := NewHandlers(associatedService)

	mux.HandleFunc("GET /ping", pingPong)

	// Associated
	mux.HandleFunc("GET /associated", handlers.AssociatedController.List())
	mux.HandleFunc("POST /associated", handlers.AssociatedController.Create())
	mux.HandleFunc("DELETE /associated/{number_card}", handlers.AssociatedController.Delete())
}
