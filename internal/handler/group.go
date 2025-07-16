package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/errs"
	"projeto-integrador-mdm/internal/service"
)

type GroupHandler struct {
	service service.GroupService
}

func NewGroupHandler(service service.GroupService) *GroupHandler {
	defer slog.Debug("criando objeto GroupHandler")
	return &GroupHandler{service: service}
}

func (h *GroupHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)

		ctx := r.Context()

		object, err := h.service.Create(ctx, r.Body)
		if err != nil {
			if errors.Is(err, errs.ErrInvalidInput) {
				slog.Error(err.Error())
				writeError(w, err.Error(), http.StatusBadRequest)
				return
			}

			if errors.Is(err, errs.ErrAlreadyExists) {
				slog.Error(err.Error())
				writeError(w, err.Error(), http.StatusBadRequest)
				return
			}

			serviceError(w, r, err)
			return
		}

		slog.Info("Registro de grupo criando")
		writeOk(w, object)
	}
}

func (h *GroupHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		id := r.PathValue("group_id")
		ctx := r.Context()

		object, err := h.service.GetById(ctx, id)
		if err != nil {
			if errors.Is(err, errs.ErrInvalidInput) {
				writeError(
					w,
					"id invalido, só é aceito conjunto de numeros.",
					http.StatusBadRequest,
				)
				return
			}

			serviceError(w, r, err)
			return
		}

		if object == nil {
			slog.Warn("Nenhum grupo encontrado com o número informado", "group_id", id)
			writeError(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
			return
		}

		slog.Info("registro de grupo encontrando", "id", object.ID)
		writeOk(w, object)
	}
}

// TODO update não esta recebndo o id pela url
func (h *GroupHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		ctx := r.Context()
		body := r.Body

		object, err := h.service.Update(ctx, body)
		if err != nil {
			if errors.Is(err, errs.ErrInvalidInput) {
				slog.Error(err.Error())
				writeError(w, err.Error(), http.StatusBadRequest)
				return
			}

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

		slog.Info("registro de grupo atualizado", "id", object.ID)
		writeOk(w, object)
	}
}

func (h *GroupHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		ctx := r.Context()

		list, err := h.service.List(ctx)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		slog.Info("Lista de grupos obtida com sucesso", "quantidade", len(list))
		writeOk(w, list)
	}
}

func (h *GroupHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)

		ctx := r.Context()
		id := r.PathValue("group_id")

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
