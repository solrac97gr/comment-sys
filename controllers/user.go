package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/solrac97gr/comments/commons"
	"github.com/solrac97gr/comments/configuration"
	"github.com/solrac97gr/comments/models"
)

// Login controller for the login
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	db := configuration.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email = ? and password=?", user.Email, pwd).First(&user)

	if ok := user.ID > 0; ok {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error: To convert the token to json %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "Usuario o Clave no válido",
			Code:    http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}
}

// UserCreate create user function
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error in read user: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	if ok := user.Password != user.ConfirmPassowrd; ok {
		m.Message = "Not same password"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)
	user.Picture = "https://gravatar.com/avatar/" + picstr + "?s=100"

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&user).Error
	if err != nil {
		m.Message = "No se puedo crear el usuario"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "User created"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)
}
