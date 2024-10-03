package handlers

import (
	"fmt"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if loggedInCustomer == nil {
		http.Error(w, "No customer is logged in", http.StatusBadRequest)
		return
	}

	logHistory(loggedInCustomer.ID, "logout", fmt.Sprintf("Customer %s logged out", loggedInCustomer.Name))
	loggedInCustomer = nil

	w.Write([]byte("Logout successful"))
}
