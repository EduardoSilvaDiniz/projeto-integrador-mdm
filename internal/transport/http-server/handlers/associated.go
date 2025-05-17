package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chamada-pagamento-system/internal/domain/entities"
	"chamada-pagamento-system/internal/migrations"
	"chamada-pagamento-system/internal/transport/http-server/dto"
)

var assocEntity entities.Associated

func getAssociated(w http.ResponseWriter, _ *http.Request) {
	var assocDtoList []dto.Associated

	if err := migrations.DB.Find(&assocDtoList).Error; err != nil {
		http.Error(w, "erro ao listar assoc: "+err.Error(), http.StatusBadRequest)
	}
	if len(assocDtoList) == 0 {
		http.Error(w, "nenhum associado encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, assocDtoList)
}

func createAssociated(w http.ResponseWriter, r *http.Request) {
	var assocDto dto.Associated

	if err := json.NewDecoder(r.Body).Decode(&assocDto); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := assocDto.IsValid(); err != nil {
		response := map[string][]string{
			"erros": err,
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, response)
		return
	}
	result := migrations.DB.Create(&assocDto)
	if err := result.Error; err != nil {
		http.Error(w, "Falha na gravação dos dados: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, assocDto)
}

func DeleteAssoc(w http.ResponseWriter, r *http.Request) {
	result := migrations.DB.Where("cpf = ?", r.PathValue("cpf")).Delete(&assocEntity)
	if err := result.Error; err != nil {
		http.Error(w, "erro ao remover associado: "+err.Error(), http.StatusBadRequest)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "nenhum associado encontrado com esse nome", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Associado deletado com sucesso")
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
		fmt.Fprintln(w, "You made a DELETE request!")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
