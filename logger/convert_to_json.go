package logger

import "encoding/json"

func ConvertToJSON(data interface{}) string {
	rv, err := json.Marshal(data)
	if err != nil {
		Error(err.Error())
	}
	return string(rv)
}
