package handler

import (
	"encoding/json"
	"net/http"

	"chamada-pagamento-system/db"
	"chamada-pagamento-system/internal/domain"
)

func createAssociated(w http.ResponseWriter, r *http.Request) {
	var assoc domain.Associated
	if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&assoc); err != nil {
		http.Error(w, "Falha ao salvar dados no banco de dados", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assoc)
}
