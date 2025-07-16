package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/errs"
	"projeto-integrador-mdm/internal/service"
)

type AssociatedHandler struct {
	service service.AssociatedService
}

func NewAssociatedHandler(service service.AssociatedService) *AssociatedHandler {
	defer slog.Debug("criando objeto AssociatedHandler")
	return &AssociatedHandler{service: service}
}

func (h *AssociatedHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)

		ctx := r.Context()

		object, err := h.service.Create(ctx, r.Body)
		if err != nil {
			if errors.Is(err, errs.ErrInvalidInput) {
				slog.Error(err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			serviceError(w, r, err)
			return
		}

		slog.Info("Registro de associado criando")
		writeOk(w, object)
	}
}

func (h *AssociatedHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		id := r.PathValue("number_card")
		ctx := r.Context()

		object, err := h.service.GetById(ctx, id)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		if object == nil {
			slog.Warn("Nenhum associado encontrado com o número informado", "number_card", id)
			writeError(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
			return
		}

		slog.Info("registro de associado encontrando", "id", object.NumberCard)
		writeOk(w, object)
	}
}

func (h *AssociatedHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		ctx := r.Context()
		body := r.Body

		object, err := h.service.Update(ctx, body)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		if object == nil {
			slog.Warn(
				"não foi encontrando registro com o numero de carterinha informado",
				"err",
				err,
			)
			writeError(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
			return
		}

		slog.Info("registro de associado atualizado", "id", object.NumberCard)
		writeOk(w, object)
	}
}

func (h *AssociatedHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		ctx := r.Context()

		list, err := h.service.List(ctx)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		slog.Info("Lista de associados obtida com sucesso", "quantidade", len(list))
		writeOk(w, list)
	}
}

func (h *AssociatedHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)

		ctx := r.Context()
		id := r.PathValue("number_card")

		rows, err := h.service.Delete(ctx, id)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		if rows == 0 {
			slog.Error("não foi encontrando registros")
			writeError(w, "Registro não encontrado", http.StatusBadRequest)
			return
		}

		slog.Info("Registro apagado", "id", id)
		writeOk(w, id)
	}
}
