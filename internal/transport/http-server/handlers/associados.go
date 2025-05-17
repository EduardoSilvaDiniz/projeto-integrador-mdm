package handlers

import (
	"chamada-pagamento-system/internal/domain/entities"
	"chamada-pagamento-system/internal/migrations"
	"encoding/json"
	"fmt"
	"net/http"
)

func getAssociated(w http.ResponseWriter, _ *http.Request) {
	var assoc []entities.Associated
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
	if r.Method != http.MethodPost {
		http.Error(w, "metodo não permitido", http.StatusMethodNotAllowed)
	}

	var assoc entities.Associated

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

func MapeamentoAssoc(w http.ResponseWriter, r *http.Request) {
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
