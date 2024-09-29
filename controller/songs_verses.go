package controller

import (
	"errors"
	"net/http"
	"songs-treasure/pkg/logging"

	"github.com/gorilla/mux"
)

type GetVersesRequest struct {
	Text string `json:"text"`
	PaginationRequest
}
type GetVersesByIdRequest struct {
	PaginationRequest
}
type EditAllVersesRequest struct {
	Id            int    `json:"id"`
	Text          string `json:"text"`
	VersePosition uint   `json:"verse_position"`
}

func (ctrl *controller) GetVerses(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("GetVerses")

	params, err := readParameters[GetVersesRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetVerses(params.Text, params.Page, params.PageSize)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) GetVersesBySongId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logging.Default.Infof("GetVersesBySongId - %s", id)
	if id == "" {
		Response(nil, errors.New("Couldn`t get request id"), w, http.StatusBadRequest)
		return
	}
	params, err := readParameters[GetVersesByIdRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetVersesBySongId(id, params.Page, params.PageSize)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) EditVerse(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("EditVerse")

	params, err := readParameters[EditAllVersesRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.EditVerse(params.Id, params.Text, params.VersePosition)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) EditAllVerses(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("EditAllVerses")

	params, err := readParameters[EditAllVersesRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.EditAllVerses(params.Id, params.Text)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}
