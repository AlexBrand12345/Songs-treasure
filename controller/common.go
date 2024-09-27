package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func Response(data interface{}, err error, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		data = err.Error()
		if status == 0 {
			status = http.StatusBadRequest
		}
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(data)
}

func readParameters[Request any](r *http.Request) (request Request, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &request)
	// if err != nil {
	// 	return
	// }
	return
}
