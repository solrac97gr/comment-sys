package models

// Message message fot the client of the API
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
