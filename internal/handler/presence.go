package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/service"
)

type PresenceHandler struct {
	service service.PresenceService
}

func NewPresenceHandler(service service.PresenceService) *PresenceHandler {
	defer slog.Debug("criando objeto PresenceHandler")
	return &PresenceHandler{service: service}
}

func (h *PresenceHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		object, err := h.service.Create(r.Context(), r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, object)
	}
}

func (h *PresenceHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		presenceList, err := h.service.List(r.Context())
		if err != nil {
			http.Error(w, "erro na execução GET: "+err.Error(), http.StatusBadRequest)
			return
		}

		if len(presenceList) == 0 {
			http.Error(w, "nenhum registro encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		for _, assoc := range presenceList {
			fmt.Fprintf(w, "%+v\n", assoc)
		}
	}
}

func (h *PresenceHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		object, err := h.service.Delete(r.Context(), r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "registro deletado com sucesso ", object)
	}
}
