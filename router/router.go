package router

import (
	"net/http"
	"songs-treasure/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Router(controller controller.Controller) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/getAuthors", controller.GetList).Methods("GET")
	// router.HandleFunc("/getSongs", controller.GetList).Methods("GET")
	// router.HandleFunc("/getCompatibility", controller.GetList).Methods("POST")
	// router.HandleFunc("/getSearchConfig", controller.GetList).Methods("GET")
	// router.HandleFunc("/setSearchConfig", controller.GetList).Methods("PUT")

	handler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Token", "Content-Type", "Origin", "Accept"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	return handler
}
