package handlers

import (
	"auth-service/app/adapter/rest/utils"
	"net/http"
)

func (h *Handlers) Hello(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, http.StatusOK, map[string]any{"success": true})
}
