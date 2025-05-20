package httpserver

import (
	"chamada-pagamento-system/internal/domain/services"
	"chamada-pagamento-system/internal/infra/repositories"
	"chamada-pagamento-system/internal/transport/http-server/handlers"
	"net/http"
)

func pingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func RegisterHandlers(mux *http.ServeMux) {
	db := repositories.PostgresMigrate()
	repo := repositories.NewGormAssociatedRepository(db)
	service := services.NewAssociatedService(repo)

	mux.HandleFunc("GET /ping", pingPong)
	// mux.HandleFunc("GET /associated", handlers.GetAllAssociatedHandler(service))
	mux.HandleFunc("POST /associated", handlers.CreateAssociatedHandler(service))

	// mux.HandleFunc("PUT /associated", handlers.MapEndpointsToAssoc)
	// mux.HandleFunc("DELETE /associated/{cpf}", handlers.DeleteAssoc)
}
