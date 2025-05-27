package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"chamada-pagamento-system/internal/transport/http-server/dto"
)

const (
	url = "http://localhost:8080/associated"
)

func TestCreateAssoc(t *testing.T) {
	assoc := dto.Associated{
		CPF:           "123",
		Name:          "edu",
		DateBirth:     "12/12/24",
		MaritalStatus: dto.MaritalStatus("Single"),
	}

	body, err := json.Marshal(assoc)
	if err != nil {
		t.Errorf("erro ao gerar JSON: %s", err.Error())
		return
	}

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		t.Fatalf("erro ao enviar POST: %s", err.Error())
	}

	defer resp.Body.Close()
	t.Logf("status: %s", resp.Status)
}

func TestGetAssoc(t *testing.T) {
	if resp, err := http.Get(url); err != nil {
		t.Fatal("erro ao receber json: ", err.Error())
	} else {
		t.Log("Status: ", resp.Status)
	}
}

func TestRemoveAssoc(t *testing.T) {
	req, err := http.NewRequest("DELETE", url+"/123", nil)
	if err != nil {
		t.Fatal("erro ao tentar criar request")
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("erro ao tentar executar resquest")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status esperado 200, recebido: ", resp.StatusCode)
	}
}
