package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chamada-pagamento-system/internal/domain/entities"
	"chamada-pagamento-system/internal/migrations"
	"chamada-pagamento-system/internal/transport/http-server/dto"
)

var assocEntitie entities.Associated

func getAssociated(w http.ResponseWriter, _ *http.Request) {
	var assoc []dto.Associated

	if err := migrations.DB.Find(&assoc).Error; err != nil {
		http.Error(w, "erro ao listar assoc: "+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")

	if len(assoc) == 0 {
		http.Error(w, "nenhum associado encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(assoc)
}

func createAssociated(w http.ResponseWriter, r *http.Request) {
	var assoc dto.Associated

	if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := assoc.IsValid(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		resposta := map[string][]string{
			"erros": err,
		}

		json.NewEncoder(w).Encode(resposta)

		return
	}

	if err := migrations.DB.Create(&assoc).Error; err != nil {
		http.Error(
			w,
			"Falha ao salvar dados no banco de dados: "+err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assoc)
}

func deleteAssoc(w http.ResponseWriter, r *http.Request) {
	var assoc dto.Associated
	if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
		http.Error(w, "erro ao pega o objeto"+err.Error(), http.StatusBadRequest)
		return
	}

	result := migrations.DB.Where("name = ?", assoc.Name).Delete(&assocEntitie)

	if err := result.Error; err != nil {
		http.Error(w, "erro ao remover "+assoc.Name+": "+err.Error(), http.StatusBadRequest)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "nenhum associado encontrado com esse nome", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Associado deletado com sucesso"))
}

func MapEndpointsToAssoc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAssociated(w, r)
	case http.MethodPost:
		createAssociated(w, r)
	case http.MethodPut:
		fmt.Fprintln(w, "You made a PUT request!")
	case http.MethodDelete:
		deleteAssoc(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
