package handler

import (
	"net/http"
	"projeto-integrador-mdm/internal/service"
)

type PresenceHandler struct {
	service service.PresenceService
}

func NewPresenceHandler(service service.PresenceService) *PresenceHandler {
	return &PresenceHandler{
		service: service,
	}
}

func (h *PresenceHandler) Create() http.HandlerFunc {
	return nil
}
