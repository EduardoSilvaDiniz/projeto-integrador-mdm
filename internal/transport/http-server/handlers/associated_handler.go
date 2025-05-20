package handlers

import (
	"chamada-pagamento-system/internal/domain/entities"
	"chamada-pagamento-system/internal/domain/services"
	"chamada-pagamento-system/internal/transport/http-server/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

var assocEntity entities.Associated

// func GetAllAssociatedHandler(svc *services.AssociatedService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, _ *http.Request) {
// 		assocList, err := svc.
// 		if err != nil {
// 			http.Error(w, "erro ao listar assoc: "+err.Error(), http.StatusBadRequest)
// 		}
//
// 		if len(assocList) == 0 {
// 			http.Error(w, "nenhum associado encontrado", http.StatusNotFound)
// 			return
// 		}
//
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, assocList)
// 	}
// }

func CreateAssociatedHandler(svc *services.AssociatedService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a dto.Associated

		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := svc.Create(&a); err != nil {
			response := map[string]string{
				"error": err.Error(),
			}

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, a)
	}
}

// func DeleteAssoc(w http.ResponseWriter, r *http.Request) {
// 	result := migrations.DB.Where("cpf = ?", r.PathValue("cpf")).Delete(&assocEntity)
// 	if err := result.Error; err != nil {
// 		http.Error(w, "erro ao remover associado: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	if result.RowsAffected == 0 {
// 		http.Error(w, "nenhum associado encontrado com esse nome", http.StatusNotFound)
// 		return
// 	}
//
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Associado deletado com sucesso")
// }
