package controller

import (
	"errors"
	"net/http"
	"songs-treasure/pkg/logging"

	"github.com/gorilla/mux"
)

type AddSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
type GetSongsRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
	Link  string `json:"link"`
	From  string `json:"from"`
	To    string `json:"to"`
	PaginationRequest
}
type GetSongsByGroupIdRequest struct {
	Song string `json:"song"`
	Link string `json:"link"`
	From string `json:"from"`
	To   string `json:"to"`
	PaginationRequest
}
type EditSongRequest struct {
	Id      int    `json:"id"`
	GroupId int    `json:"group_id"`
	Song    string `json:"song"`
	Link    string `json:"link"`
	Release string `json:"release_date"`
}
type DeleteSongRequest struct {
	Id int `json:"id"`
}

func (ctrl *controller) AddSong(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("AddSong")

	params, err := readParameters[AddSongRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.AddSong(params.Group, params.Song)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) GetSong(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logging.Default.Infof("GetSong - %s", id)
	if id == "" {
		Response(nil, errors.New("Couldn`t get request id"), w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetSong(id)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) GetSongs(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("GetSongs")

	params, err := readParameters[GetSongsRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetSongs(params.Group,
		params.Song, params.From, params.To, params.Link, params.Page, params.PageSize)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) GetSongsByGroupId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logging.Default.Infof("GetSongsByGroupId - %s", id)
	if id == "" {
		Response(nil, errors.New("Couldn`t get request id"), w, http.StatusBadRequest)
		return
	}

	params, err := readParameters[GetSongsByGroupIdRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.GetSongsByGroupId(id,
		params.Song, params.From, params.To, params.Link, params.Page, params.PageSize)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) EditSong(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("EditSong")

	params, err := readParameters[EditSongRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	data, err := ctrl.service.EditSong(params.Id,
		params.GroupId, params.Song, params.Release, params.Link)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(data, err, w, http.StatusOK)
}

func (ctrl *controller) DeleteSong(w http.ResponseWriter, r *http.Request) {
	logging.Default.Info("DeleteSong")

	params, err := readParameters[DeleteSongRequest](r)
	if err != nil {
		Response(nil, err, w, http.StatusBadRequest)
		return
	}

	err = ctrl.service.DeleteSong(params.Id)
	if err != nil {
		Response(nil, err, w, http.StatusInternalServerError)
		return
	}

	Response(nil, err, w, http.StatusOK)
}
