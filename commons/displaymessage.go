package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/solrac97gr/comments/models"
)

// DisplayMessage return message to the client
func DisplayMessage(w http.ResponseWriter, message models.Message) {
	j, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error to conver the message: %s", err)
	}
	w.WriteHeader(message.Code)
	w.Write(j)
}
