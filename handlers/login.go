package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"first-go-project/models"
)

var loggedInCustomer *models.Customer = nil

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	var customer models.Customer
	query := "SELECT id, name, email, password FROM customers WHERE email=$1 AND password=$2"

	err := DB.QueryRow(query, email, password).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error during login", http.StatusInternalServerError)
		return
	}

	loggedInCustomer = &customer
	logHistory(customer.ID, "login", fmt.Sprintf("Customer %s logged in", customer.Name))

	w.Write([]byte("Login successful"))
}
