package handlers

import (
	"auth-service/app/adapter/rest/utils"
	"auth-service/types"
	"net/http"
)

func (h *Handlers) SignInUser(w http.ResponseWriter, r *http.Request) {
	var payload types.SignInUserPayload

	// decode json request body
	err := utils.DecodeJSON(r, &payload)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	err = payload.Validate()
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	accessToken, refreshToken, err := h.userService.LoginUser(r.Context(), payload)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to login user", err.Error())
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
