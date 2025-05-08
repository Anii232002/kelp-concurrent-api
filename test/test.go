package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type testRequest struct {
	CompanyID string
	Endpoint  string
	Delay     time.Duration
}

func main() {
	var wg sync.WaitGroup

	requests := []testRequest{
		{"123", "/api/financials/data", 0 * time.Millisecond},
		{"123", "/api/sales/data", 100 * time.Millisecond},
		{"123", "/api/employee", 0 * time.Millisecond},
		{"123", "/api/financials/data", 5 * time.Second},
		{"331", "/api/sales/data", 0 * time.Millisecond},
		{"331", "/api/sales/data", 5 * time.Second},
	}

	for _, req := range requests {
		wg.Add(1)
		go func(companyID string, endpoint string, delay time.Duration) {
			defer wg.Done()
			time.Sleep(delay)

			url := fmt.Sprintf("http://localhost:8080%s?companyId=%s", endpoint, companyID)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("error %s %s -> %v\n", companyID, endpoint, err)
				return
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)

			fmt.Printf("http %s %s \n\t-> %s\n", companyID, endpoint, string(body))
		}(req.CompanyID, req.Endpoint, req.Delay)
	}

	wg.Wait()
	fmt.Println("All HTTP calls completed.")
}
