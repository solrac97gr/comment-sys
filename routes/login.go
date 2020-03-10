package routes

import (
	"github.com/gorilla/mux"
	"github.com/solrac97gr/comments/controllers"
)

// SetLoginRouter Seting the routes for login control
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
