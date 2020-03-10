package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/solrac97gr/comments/commons"
	"github.com/solrac97gr/comments/configuration"
	"github.com/solrac97gr/comments/models"
)

// CommentCreate Create a comment in the DB
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error when read the comment: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error when register the comment: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comment created"
	commons.DisplayMessage(w, m)
}

// GetComments Return all comments of the db
func GetComments(w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	m := models.Message{}

}
