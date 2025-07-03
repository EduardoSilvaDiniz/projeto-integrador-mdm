package endpoint

import (
	"chamada-pagamento-system/internal/database"
	"chamada-pagamento-system/internal/endpoint/handler"
	"chamada-pagamento-system/internal/infra/repositories"
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

func NewHandlers(assoc handler.AssociatedController) *Handlers {
	return &Handlers{AssociatedController: assoc}
}

func CreateEndpoints(mux *http.ServeMux) {
	conn := repositories.PgxConnect()
	queries := database.New(conn)
	associatedService := handler.NewAssociatedService(queries)
	handlers := NewHandlers(associatedService)

	mux.HandleFunc("GET /ping", pingPong)

	// Associated
	mux.HandleFunc("GET /associated", handlers.AssociatedController.List())
	mux.HandleFunc("POST /associated", handlers.AssociatedController.Create())
	// mux.HandleFunc("PUT /associated", handlers.MapEndpointsToAssoc)
	mux.HandleFunc("DELETE /associated/{cpf}", handlers.AssociatedController.Delete())
}
