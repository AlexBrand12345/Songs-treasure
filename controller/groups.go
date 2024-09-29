package controller

import (
	"errors"
	"net/http"
	"songs-treasure/pkg/logging"

	"github.com/gorilla/mux"
)

type GetGroupsRequest struct {
	Group string `json:"group"` // group name
	PaginationRequest
}

func (ctrl *controller) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logging.Default.Infof("GetGroup - %s", id)
	if id == "" {
		Response(nil, errors.New("Couldn`t get request id"), w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetGroup(id)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) GetGroups(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("GetGroups")

	params, err := readParameters[GetGroupsRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetGroups(params.Group, params.Page, params.PageSize)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}
