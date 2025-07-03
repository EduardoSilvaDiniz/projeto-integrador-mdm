package endpoint

import (
	"chamada-pagamento-system/internal/database"
	"chamada-pagamento-system/internal/endpoint/handler/associated"
	"chamada-pagamento-system/internal/infra/repositories"
	"net/http"
)

func pingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateEndpoints(mux *http.ServeMux) {
	conn := repositories.PgxConnect()
	queries := database.New(conn)
	associatedService := associated.NewAssociatedService(queries)
	handlers := associated.NewHandlers(associatedService)

	// db := repositories.PostgresMigrate()
	// repo := repositories.NewGormAssociatedRepository(db)
	// service := services.NewAssociatedService(repo)

	mux.HandleFunc("GET /ping", pingPong)
	mux.HandleFunc("GET /associated", handlers.AssociatedController.List())
	mux.HandleFunc("POST /associated", handlers.AssociatedController.Create())

	// mux.HandleFunc("PUT /associated", handlers.MapEndpointsToAssoc)
	// mux.HandleFunc("DELETE /associated/{cpf}", handlers.DeleteAssoc)
}
