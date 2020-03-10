package routes

import "github.com/gorilla/mux"

// InitRoutes Init the routes and return the router
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)

	return router
}
