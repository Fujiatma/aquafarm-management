package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func ConstructResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	response := Response{}

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()

		status = http.StatusInternalServerError
	} else {
		response.Data = data
		response.Status = "success"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
