package handlers

import (
	"auth-service/app/adapter/rest/utils"
	"auth-service/types"

	"net/http"
)

func (h *Handlers) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var payload types.SignUpUserPayload

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

	err = h.userService.Create(r.Context(), payload)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "failed to create user", err.Error())
		return
	}

	utils.SendData(w, http.StatusCreated)

}
