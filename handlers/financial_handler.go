package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Anii232002/concurrent-api/calculations"
)

func FinancialHandler(w http.ResponseWriter, r *http.Request) {
	companyId := r.URL.Query().Get("companyId")

	if companyId == "" {
		http.Error(w, "companyId is required", http.StatusBadRequest)
		return
	}
	fmt.Printf("new api call , company : %s , url : %s\n", companyId, r.URL)

	result := calculations.ProcessRequest(companyId, "financials")

	response := map[string]int{
		"result": result,
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
