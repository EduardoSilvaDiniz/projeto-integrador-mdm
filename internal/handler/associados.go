package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"chamada-pagamento-system/internal/domain"
)

func createAssociated(w http.ResponseWriter, r *http.Request) {
	var assoc domain.Associated
	if err := json.NewDecoder(r.Body).Decode(&assoc); err != nil {
		http.Error(w, "JSON invalido: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Associado recebido: ", assoc)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"message": "associado criando com sucesso",
	}); err != nil {
		log.Println("Erro ao escrever resposta: ", err)
	}
}
