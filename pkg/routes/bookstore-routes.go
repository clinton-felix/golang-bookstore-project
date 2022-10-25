package routes

import (
	"github.com/clinton-felix/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// Creating a Register Book store Router function
// This function creates the handlers for the different routes
var RegisterBookStoreRoutes = func(router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}