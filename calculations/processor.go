package calculations

import (
	"fmt"
	"math/rand"
	"time"
)

func ProcessRequest(companyId string, apiName string) int {

	key := fmt.Sprintf("%s:%s", companyId, apiName)

	if data, exists := checkCacheWithLog(key, apiResultCache, "api result"); exists {
		return data
	}

	if data, exists := waitForConcurrentRequest(key, pendingRequests, "api result"); exists {
		return data
	}

	return startProcessingRequest(key, companyId, apiName)
}

func startProcessingRequest(key string, companyId string, apiName string) int {

	waitCh := make(chan int, 1)
	mu.Lock()
	pendingRequests[key] = &PendingRequest{
		waitCh: waitCh,
	}
	mu.Unlock()

	initialData := CalculateInitialData(companyId)

	fmt.Printf("processing api '%s' for company: %s\n", apiName, companyId)
	time.Sleep(2 * time.Second)

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	result := initialData*randomNumber + 1000

	waitCh <- result
	mu.Lock()
	apiResultCache[key] = result
	delete(pendingRequests, key)
	mu.Unlock()

	close(waitCh)
	return result
}
