package utils

import (
	"auth-service/logger"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	str, err := json.Marshal(data)

	if err != nil {
		slog.Error(err.Error(), logger.Extra(map[string]any{
			"data": err.Error(),
		}))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)
	w.Write(str)
}
