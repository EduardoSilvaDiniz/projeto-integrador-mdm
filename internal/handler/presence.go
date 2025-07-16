
package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/errs"
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

		slog.Info("Registro de presença criando")
		writeOk(w, object)
	}
}

func (h *PresenceHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		numberCard := r.PathValue("number_card")
		meetingId := r.PathValue("meeting_id")
		ctx := r.Context()

		object, err := h.service.GetById(ctx, numberCard, meetingId)
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
			slog.Warn("Nenhum associado encontrado com o número informado", "number_card", numberCard, "meeting_id", meetingId)
			writeError(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
			return
		}

		slog.Info("registro de associado encontrando", "id", object)
		writeOk(w, object)
	}
}

func (h *PresenceHandler) Update() http.HandlerFunc {
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

		slog.Info("registro de presença atualizado", "id", object.NumberCard)
		writeOk(w, object)
	}
}

func (h *PresenceHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		ctx := r.Context()

		list, err := h.service.List(ctx)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		slog.Info("Lista de presenças obtida com sucesso", "quantidade", len(list))
		writeOk(w, list)
	}
}

func (h *PresenceHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)

		ctx := r.Context()

		object, err := h.service.Delete(ctx, r.Body)
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

		slog.Info("Registro apagado", "id", object)
		writeOk(w, object)
	}
}
