package controller

import (
	"songs-treasure/internal/service"
)

type Controller interface {
	// GetAuthors(w http.ResponseWriter, r *http.Request)
	// GetSongsWithAutors(w http.ResponseWriter, r *http.Request)

	// GetSong(w http.ResponseWriter, r *http.Request)
	// DeleteSong(w http.ResponseWriter, r *http.Request)
	// UpdateSong(w http.ResponseWriter, r *http.Request)
	// CreateSong(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service *service.Service
}

func NewController(service *service.Service) *controller {
	return &controller{
		service: service,
	}
}
