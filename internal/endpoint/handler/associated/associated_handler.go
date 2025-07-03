package associated

import (
	"chamada-pagamento-system/internal/database"
	"fmt"
	"net/http"
)
type AssociatedController interface {
	// Create() error
	// search() error
	List() http.HandlerFunc
	// delete() error
}

type AssociatedService struct{
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
			http.Error(w, "erro ao listar assoc: " + err.Error(), http.StatusBadRequest)
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

//
// func CreateAssociatedHandler(svc *services.AssociatedService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var assoc dto.Associated
//
// 		if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
// 			http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
// 			return
// 		}
//
// 		if err := svc.Create(&assoc); err != nil {
// 			response := map[string]string{
// 				"error": err.Error(),
// 			}
//
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, response)
// 			return
// 		}
//
// 		w.WriteHeader(http.StatusCreated)
// 		fmt.Fprintln(w, assoc)
// 	}
// }

// func DeleteAssocHandler(svc *services.AssociatedService) http.HandlerFunc {
// return func(w http.ResponseWriter, r *http.Request) {
// 	if err := svc.Repo.DeleteByCPF(r.PathValue("cpf")); err != nil {
// 		http.Error(w, "erro ao remover associado: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Associado deletado com sucesso")
// }
// }
