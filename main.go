package main

import (
	"log"
	"net/http"

	"github.com/Anii232002/concurrent-api/handlers"
)

func main() {
	http.HandleFunc("/api/financials/data", handlers.FinancialHandler)
	http.HandleFunc("/api/sales/data", handlers.SalesHandler)
	http.HandleFunc("/api/employee", handlers.EmployeeHandler)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
