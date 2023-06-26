package routes

import (
	"github.com/Manda-supraja26/moviestore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{movieid}", controllers.GetMovieByID).Methods("Get")
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{movieid}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieid}", controllers.DeleteMovie).Methods("DELETE")

}
