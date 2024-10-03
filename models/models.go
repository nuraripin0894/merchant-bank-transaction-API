package models

type History struct {
	CustomerID  int    `json:"customer_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

type Customer struct {
	ID       int
	Name     string
	Email    string
	Password string
}
