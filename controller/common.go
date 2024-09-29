package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"songs-treasure/pkg/logging"
)

type PaginationRequest struct {
	PageSize uint `json:"page_size"`
	Page     uint `json:"page"`
}

func Response(data interface{}, err error, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		data = err.Error()
		if status == 0 {
			status = http.StatusBadRequest
		}
		w.WriteHeader(status)
	} else {
		logging.Default.Info(fmt.Sprintf("%+v", data))
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(data)
}

func readParameters[Request any](r *http.Request) (request Request, err error) {
	defer func() {
		if err != nil {
			logging.Default.Error(err.Error())
		} else {
			logging.Default.Info(fmt.Sprintf("%+v", request))
		}
	}()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &request)

	return
}
