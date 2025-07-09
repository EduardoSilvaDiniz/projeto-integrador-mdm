package handler

import (
	"chamada-pagamento-system/internal/service"
	"fmt"
	"net/http"
)

type AssociatedHandler struct {
	service service.AssociatedService
}

func NewAssociatedHandler(service service.AssociatedService) *AssociatedHandler {
	return &AssociatedHandler{
		service: service,
	}
}

func (h *AssociatedHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		associatedList, err := h.service.List(r.Context())

		if err != nil {
			http.Error(w, "erro na execução GetAssociated: "+err.Error(), http.StatusBadRequest)
			return
		}

		if len(associatedList) == 0 {
			http.Error(w, "nenhum associado encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		for _, assoc := range associatedList {
			fmt.Fprintf(w, "%+v\n", assoc)
		}
	}
}

func (h *AssociatedHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		associatedObject, err := h.service.Create(r.Context(), r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, associatedObject)
	}
}

func (h *AssociatedHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cardNumber := r.PathValue("number_card")
		rows, err := h.service.Delete(ctx, cardNumber)
		if err != nil {
			http.Error(w, "ocorreu um erro no service", http.StatusInternalServerError)
			return
		}

		if rows == 0 {
			http.Error(w, "Registro não encontrado", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Associado deletado com sucesso")
	}
}
