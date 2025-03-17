package handlers

import (
	"auth-service/app/adapter/rest/utils"
	"auth-service/types"

	"net/http"
)

func (h *Handlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var payload types.RefreshTokenPayload

	// decode json request body
	err := utils.DecodeJSON(r, &payload)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	claims, err := h.userService.ValidateToken(r.Context(), payload.RefreshToken)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to validate token", err.Error())
		return
	}

	if claims.Type != "refresh" {
		utils.SendError(w, http.StatusBadRequest, "Invalid token type", err)
		return
	}

	userID := claims.ID
	user, err := h.userService.GetUserById(r.Context(), userID)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to get user profile", err.Error())
		return
	}

	accessToken, err := h.userService.GenerateToken(user, h.conf.JWT.AccessExpIn, "access")

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to generate token", err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]any{
		"accessToken": accessToken,
	})

}
