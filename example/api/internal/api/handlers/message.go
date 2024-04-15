package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Assuming you want a simple in-memory store
var messages = make(map[string]string)

// Message structure
type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// CreateMessage - POST /message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Store the message
	messages[msg.ID] = msg.Content
	json.NewEncoder(w).Encode(msg)
}

// GetMessage - GET /message/{id}
func GetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	content, ok := messages[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(Message{ID: id, Content: content})
}
