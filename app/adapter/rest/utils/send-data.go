package utils

import (
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}) {
	SendJSON(w, http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Success",
		"data":    data,
	})
}
