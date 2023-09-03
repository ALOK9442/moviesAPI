package router

import (
	"fmt"

	"github.com/ALOK9442/mongoAPI/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	fmt.Println("welcome to database")
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeletedAllMovies).Methods("DELETE")

	return router
}
