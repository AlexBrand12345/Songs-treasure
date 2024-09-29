package router

import (
	"net/http"
	"songs-treasure/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Router(controller controller.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/getGroup/{id}", controller.GetGroup).Methods("GET")
	router.HandleFunc("/getGroups", controller.GetGroups).Methods("GET")

	router.HandleFunc("/getSong/{id}", controller.GetSong).Methods("GET")
	router.HandleFunc("/getSongs", controller.GetSongs).Methods("GET")
	router.HandleFunc("/getSongs/{id}", controller.GetSong).Methods("GET")
	router.HandleFunc("/addSong", controller.AddSong).Methods("POST")
	router.HandleFunc("/editSong", controller.EditSong).Methods("PUT")
	router.HandleFunc("/deleteSong", controller.DeleteSong).Methods("DELETE")

	router.HandleFunc("/getVerses", controller.GetVerses).Methods("GET")
	router.HandleFunc("/getVerses/{id}", controller.GetVersesBySongId).Methods("GET")
	router.HandleFunc("/editVerse", controller.EditVerse).Methods("PATCH")
	router.HandleFunc("/editAllVerses", controller.EditAllVerses).Methods("PUT")

	handler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Token", "Content-Type", "Origin", "Accept"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	return handler
}
