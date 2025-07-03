package associated

import (
	"chamada-pagamento-system/internal/database"
	"chamada-pagamento-system/internal/domain/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

type AssociatedController interface {
	Create() http.HandlerFunc
	// search() error
	List() http.HandlerFunc
	Delete() http.HandlerFunc
}

type AssociatedService struct {
	queries *database.Queries
}

func NewAssociatedService(q *database.Queries) *AssociatedService {
	return &AssociatedService{
		queries: q,
	}
}

func (a *AssociatedService) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		assocList, err := a.queries.GetAssoc(ctx)

		if err != nil {
			http.Error(w, "erro ao listar assoc: "+err.Error(), http.StatusBadRequest)
		}

		if len(assocList) == 0 {
			http.Error(w, "nenhum associado encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		for _, assoc := range assocList {
			fmt.Fprintf(w, "%+v\n", assoc)
		}
	}
}

func (a *AssociatedService) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var assoc entity.Associated

		if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
			http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
			return
		}

		assocParam := database.CreateAssocParams{
			Cpf:           assoc.CPF,
			Name:          assoc.Name,
			DateBirth:     assoc.DateBirth,
			MaritalStatus: string(assoc.MaritalStatus),
		}

		if err := a.queries.CreateAssoc(ctx, assocParam); err != nil {
			response := map[string]string{
				"error": err.Error(),
			}

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, assoc)
	}
}

func (a *AssociatedService) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := a.queries.DeleteAssoc(ctx, r.PathValue("cpf")); err != nil {
			http.Error(w, "erro ao remover associado: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Associado deletado com sucesso")
	}
}
