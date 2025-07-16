package web

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/db"
	"projeto-integrador-mdm/internal/handler"
	"projeto-integrador-mdm/internal/service"
)

func PingPong(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	ua := r.UserAgent()
	method := r.Method
	path := r.URL.Path
	msg := "pong"

	slog.Info("Requisição recebida",
		"ip", ip,
		"user_agent", ua,
		"method", method,
		"path", path,
	)

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		slog.Error("erro ao tentar enviar JSON", "err", err)
		http.Error(w, "erro ao tentar enviar JSON", http.StatusInternalServerError)
	}
}

func CreateRouter(mux *http.ServeMux, queries *db.Queries) {
	defer slog.Debug("endpoints criados")

	associatedService := service.NewAssociatedService(queries)
	associatedHandler := handler.NewAssociatedHandler(associatedService)
	presenceService := service.NewPresenceService(queries)
	presenceHandler := handler.NewPresenceHandler(presenceService)
	paymentService := service.NewPaymentService(queries)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	mux.HandleFunc("GET /ping", PingPong)

	mux.HandleFunc("GET /associated", associatedHandler.List())
	mux.HandleFunc("GET /associated/{number_card}", associatedHandler.GetById())
	mux.HandleFunc("PUT /associated", associatedHandler.Update())
	mux.HandleFunc("POST /associated", associatedHandler.Create())
	mux.HandleFunc("DELETE /associated/{number_card}", associatedHandler.Delete())

	mux.HandleFunc("GET /presence", presenceHandler.List())
	mux.HandleFunc("POST /presence", presenceHandler.Create())
	mux.HandleFunc("DELETE /presence", presenceHandler.Delete())

	mux.HandleFunc("GET /payment", paymentHandler.List())
	mux.HandleFunc("POST /payment", paymentHandler.Create())
	mux.HandleFunc("DELETE /payment/{payment_id}", paymentHandler.Delete())
}
