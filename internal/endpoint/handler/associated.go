package handler

import (
	"chamada-pagamento-system/internal/database"
	"chamada-pagamento-system/internal/domain/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type AssociatedController interface {
	Create() http.HandlerFunc
	List() http.HandlerFunc
	Delete() http.HandlerFunc
}

type AssociatedService struct {
	repo *database.Queries
}

func NewAssociatedService(queries *database.Queries) *AssociatedService {
	return &AssociatedService{repo: queries}
}

func (a *AssociatedService) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		associatedList, err := a.repo.GetAssociated(ctx)

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

func (a *AssociatedService) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var assoc entity.Associated

		if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
			http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
			return
		}

		associatedParam := database.CreateAssociatedParams{
			NumberCard: assoc.NumberCard,
			Name:       assoc.Name,
		}

		if err := a.repo.CreateAssociated(ctx, associatedParam); err != nil {
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
		getCode := r.PathValue("number_card")
		conv, err := strconv.ParseInt(getCode, 10, 32)

		if err != nil {
			log.Fatalln(err)
		}

		result, err := a.repo.DeleteAssociatedByNumberCard(ctx, conv)
		if err != nil {
			http.Error(w, "erro ao remover associado: "+err.Error(), http.StatusBadRequest)
			return
		}

		rows, err := result.RowsAffected()
		if err != nil {
			log.Println("Erro ao verificar rows affected:", err)
			http.Error(w, "Erro interno", http.StatusInternalServerError)
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
