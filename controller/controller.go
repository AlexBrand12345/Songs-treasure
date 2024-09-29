package controller

import (
	"net/http"
	"songs-treasure/internal/service"
)

type Controller interface {
	GetGroup(w http.ResponseWriter, r *http.Request) // by id
	GetGroups(w http.ResponseWriter, r *http.Request)

	AddSong(w http.ResponseWriter, r *http.Request)
	GetSong(w http.ResponseWriter, r *http.Request) // by id
	GetSongs(w http.ResponseWriter, r *http.Request)
	GetSongsByGroupId(w http.ResponseWriter, r *http.Request)
	EditSong(w http.ResponseWriter, r *http.Request)
	DeleteSong(w http.ResponseWriter, r *http.Request)

	GetVerses(w http.ResponseWriter, r *http.Request)
	GetVersesBySongId(w http.ResponseWriter, r *http.Request)
	EditVerse(w http.ResponseWriter, r *http.Request)
	EditAllVerses(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{
		service: service,
	}
}
