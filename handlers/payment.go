package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if loggedInCustomer == nil {
		http.Error(w, "Please log in first", http.StatusUnauthorized)
		return
	}

	merchantIDStr := r.FormValue("merchant_id")
	amountStr := r.FormValue("amount")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil {
		http.Error(w, "Invalid merchant_id", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Check if the merchant exists
	var exists bool
	err = DB.QueryRow("SELECT EXISTS (SELECT 1 FROM merchants WHERE id=$1)", merchantID).Scan(&exists)
	if err != nil {
		http.Error(w, "Error checking merchant existence", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Merchant not found", http.StatusBadRequest)
		return
	}

	status := "COMPLETED"

	// Insert transaction into the transactions table
	_, err = DB.Exec("INSERT INTO transactions (customer_id, merchant_id, amount, status) VALUES ($1, $2, $3, $4)", loggedInCustomer.ID, merchantID, amount, status)
	if err != nil {
		log.Printf("Error processing transaction: %v", err)
		http.Error(w, "Error processing transaction", http.StatusInternalServerError)
		return
	}

	logHistory(loggedInCustomer.ID, "payment", fmt.Sprintf("Customer %s paid %.2f to Merchant %d", loggedInCustomer.Name, amount, merchantID))
	w.Write([]byte(fmt.Sprintf("Payment of %.2f successful to merchant %d", amount, merchantID)))
}
