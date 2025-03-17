package handlers

import (
	"auth-service/app/adapter/rest/middlewares"
	"auth-service/app/adapter/rest/utils"
	"net/http"
)

func (h *Handlers) GetUserDetails(w http.ResponseWriter, r *http.Request) {

	user, err := r.Context().Value(middlewares.UserKey).(*middlewares.AuthClaims)

	if !err || user == nil {
		utils.SendJSON(w, http.StatusUnauthorized, map[string]any{
			"success": false,
			"message": "Unauthorized: User not found in context",
		})
		return
	}

	userResponse, ok := h.userService.GetProfile(r.Context(), user.ID)

	if ok != nil {
		utils.SendJSON(w, http.StatusInternalServerError, map[string]any{
			"success": false,
			"message": "Failed to get user details",
		})
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    userResponse,
	})
}
