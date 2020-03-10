package routes

import (
	"github.com/gorilla/mux"
	"github.com/solrac97gr/comments/controllers"
)

func SetCommentRouter(router *mux.Router) {
	router.HandleFunc("/api/comments", controllers.CommentCreate).Methods("POST")
}
