package handler

import (
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/ping", pingPong)
	mux.HandleFunc("/associated", createAssociated)
}

func pingPong(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
