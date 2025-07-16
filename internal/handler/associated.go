package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/service"
)

type AssociatedHandler struct {
	service service.AssociatedService
}

func NewAssociatedHandler(service service.AssociatedService) *AssociatedHandler {
	defer slog.Debug("criando objeto AssociatedHandler")
	return &AssociatedHandler{
		service: service,
	}
}

func (h *AssociatedHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path
		ctx := r.Context()

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		object, err := h.service.Create(ctx, r.Body)
		if err != nil {
			slog.Error("Erro ao criar registro de associados", "err", err)
			http.Error(
				w,
				"erro ao tentar criar registro de associado",
				http.StatusInternalServerError,
			)
			return
		}
		slog.Info("Registro de associado criando")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(object); err != nil {
			slog.Error("erro ao tentar enviar JSON", "err", err)
			http.Error(w, "erro ao tentar enviar JSON", http.StatusInternalServerError)
		}
	}
}

func (h *AssociatedHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path
		id := r.PathValue("number_card")
		ctx := r.Context()

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		object, err := h.service.GetById(ctx, id)
		if err != nil {
			slog.Error("erro ao tentar busca registro de associado", "err", err)
			http.Error(
				w,
				"erro ao tentar busca registro de associado",
				http.StatusInternalServerError,
			)
			return
		}
		if object == nil {
			slog.Error(
				"não foi encontrando registro com o numero de carterinha informado",
				"err",
				err,
			)
			http.Error(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
		}

		slog.Info("registro de associado encontrando", "id", object.NumberCard)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(object); err != nil {
			slog.Error("erro ao tentar enviar JSON", "err", err)
			http.Error(w, "erro ao tentar enviar JSON", http.StatusInternalServerError)
		}
	}
}

func (h *AssociatedHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path
		ctx := r.Context()
		body := r.Body

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		object, err := h.service.Update(ctx, body)
		if err != nil {
			http.Error(
				w,
				"erro ao tentar atualizar registro de associado",
				http.StatusInternalServerError,
			)
			return
		}

		if object == nil {
			slog.Error(
				"não foi encontrando registro com o numero de carterinha informado",
				"err",
				err,
			)
			http.Error(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
		}

		slog.Info("registro de associado atualizado", "id", object.NumberCard)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(object); err != nil {
			slog.Error("erro ao tentar enviar JSON", "err", err)
			http.Error(w, "erro ao tentar enviar JSON", http.StatusInternalServerError)
		}
	}
}

func (h *AssociatedHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path
		ctx := r.Context()

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		list, err := h.service.List(ctx)
		if err != nil {
			slog.Error("Erro ao buscar lista de associados", "err", err)
			http.Error(w, "erro ao buscar associados", http.StatusInternalServerError)
			return
		}

		slog.Info("Lista de associados obtida com sucesso", "quantidade", len(list))

		if len(list) == 0 {
			http.Error(w, "nenhum registro encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(list); err != nil {
			slog.Error("erro ao tentar enviar JSON", "err", err)
			http.Error(w, "erro ao tentar enviar JSON", http.StatusInternalServerError)
			return
		}
	}
}

func (h *AssociatedHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ua := r.UserAgent()
		method := r.Method
		path := r.URL.Path
		ctx := r.Context()
		id := r.PathValue("number_card")

		slog.Info("Requisição recebida",
			"ip", ip,
			"user_agent", ua,
			"method", method,
			"path", path,
		)

		rows, err := h.service.Delete(ctx, id)
		if err != nil {
			slog.Error("erro ao tentar apagar registro", "err", err)
			http.Error(w, "erro ao tentar apagar registro", http.StatusInternalServerError)
			return
		}

		if rows == 0 {
			slog.Error("não foi encontrando registros")
			http.Error(w, "Registro não encontrado", http.StatusBadRequest)
			return
		}

		slog.Info("Registro apagado", "id", id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(id); err != nil {
			slog.Error("erro ao tentar enviar JSON", "err", err)
			http.Error(w, "erro ao tentar enviar JSON", http.StatusInternalServerError)
			return
		}
	}
}
