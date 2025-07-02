package handlers

import (
	// "chamada-pagamento-system/internal/domain/services"
	// "chamada-pagamento-system/internal/transport/http-server/dto"
	// "encoding/json"
	// "fmt"
	// "net/http"
)


// func GetAllAssociatedHandler(svc *services.AssociatedService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, _ *http.Request) {
// 		assocList, err := svc.Repo.GetAll()
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
