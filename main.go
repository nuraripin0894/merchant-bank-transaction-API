package main

import (
	"log"
	"net/http"

	"first-go-project/databases"
	"first-go-project/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = databases.InitDB()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	handlers.SetDB(databases.DB)

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/payment", handlers.PaymentHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/history", handlers.HistoryHandler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
