package handlers

import (
	"encoding/json"
	"os"
	"time"

	"first-go-project/models"
)

func logHistory(customerID int, action, description string) error {
	entry := models.History{
		CustomerID:  customerID,
		Action:      action,
		Description: description,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	file, err := os.OpenFile("history.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	_, err = file.Write([]byte("\n"))
	return err
}

// Helper function to split JSON lines
func splitJSONLines(data []byte) []string {
	lines := make([]string, 0)
	line := ""
	for _, b := range data {
		line += string(b)
		if b == '\n' {
			lines = append(lines, line)
			line = ""
		}
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}
	return lines
}
