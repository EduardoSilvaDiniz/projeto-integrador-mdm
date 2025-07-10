package web

import (
	"net/http"
	"projeto-integrador-mdm/internal/database"
	"projeto-integrador-mdm/internal/handler"
	"projeto-integrador-mdm/internal/service"
)

func PingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func CreateRouter(mux *http.ServeMux, queries *database.Queries) {
	associatedService := service.NewAssociatedService(queries)
	associatedHandler := handler.NewAssociatedHandler(associatedService)
	presenceService := service.NewPresenceService(queries)
	presenceHandler := handler.NewPresenceHandler(presenceService)
	paymentService := service.NewPaymentService(queries)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	mux.HandleFunc("GET /ping", PingPong)

	mux.HandleFunc("GET /associated", associatedHandler.List())
	mux.HandleFunc("POST /associated", associatedHandler.Create())
	mux.HandleFunc("DELETE /associated/{number_card}", associatedHandler.Delete())

	mux.HandleFunc("GET /presence", presenceHandler.List())
	mux.HandleFunc("POST /presence", presenceHandler.Create())
	mux.HandleFunc("DELETE /presence", presenceHandler.Delete())

	mux.HandleFunc("GET /payment", paymentHandler.List())
	mux.HandleFunc("POST /payment", paymentHandler.Create())
	mux.HandleFunc("DELETE /payment/{payment_id}", paymentHandler.Delete())
}
