
package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"projeto-integrador-mdm/internal/errs"
	"projeto-integrador-mdm/internal/service"
)

type MeetingHandler struct {
	service service.MeetingService
}

func NewMeetingHandler(service service.MeetingService) *MeetingHandler {
	defer slog.Debug("criando objeto MeetingHandler")
	return &MeetingHandler{service: service}
}

func (h *MeetingHandler) Create() http.HandlerFunc {
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

		slog.Info("Registro de reunião criando")
		writeOk(w, object)
	}
}

func (h *MeetingHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		id := r.PathValue("meeting_id")
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
			slog.Warn("Nenhum reunião encontrado com o número informado", "meeting_id", id)
			writeError(
				w,
				"não foi encontrando registro com numero de carterinha informado",
				http.StatusBadRequest,
			)
			return
		}

		slog.Info("registro de reunião encontrando", "id", object.ID)
		writeOk(w, object)
	}
}

//TODO o update não esta usando o id passado pela url
func (h *MeetingHandler) Update() http.HandlerFunc {
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

		slog.Info("registro de reunião atualizado", "id", object.ID)
		writeOk(w, object)
	}
}

func (h *MeetingHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)
		ctx := r.Context()

		list, err := h.service.List(ctx)
		if err != nil {
			serviceError(w, r, err)
			return
		}

		slog.Info("Lista de reuniãos obtida com sucesso", "quantidade", len(list))
		writeOk(w, list)
	}
}

func (h *MeetingHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logInicial(r)

		ctx := r.Context()
		id := r.PathValue("meeting_id")

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
