package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/solrac97gr/comments/controllers"
)

// SetCommentRouter Init routes of comments
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
