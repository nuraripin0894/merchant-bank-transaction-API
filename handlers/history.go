package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"first-go-project/models"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Open the history.json file
	file, err := os.Open("history.json")
	if err != nil {
		log.Printf("Error opening history file: %v", err)
		http.Error(w, "Error reading history file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file contents", http.StatusInternalServerError)
		return
	}

	// Convert the content to History structs
	var history []models.History
	lines := splitJSONLines(content)
	for _, line := range lines {
		var entry models.History
		if err := json.Unmarshal([]byte(line), &entry); err == nil {
			history = append(history, entry)
		}
	}

	// Return the history as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
