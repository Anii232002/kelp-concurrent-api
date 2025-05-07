package calculations

import (
	"fmt"
	"strconv"
	"time"
)

func CalculateInitialData(companyId string) int {

	if data, found := checkCacheWithLog(companyId, initialDataCache, "initial data"); found {
		return data
	}

	if data, found := waitForConcurrentRequest(companyId, initialDataPending, "initial data"); found {
		return data
	}

	return startInitialDataCalculation(companyId)

}

func startInitialDataCalculation(companyId string) int {
	waitCh := make(chan int, 1)

	mu.Lock()
	initialDataPending[companyId] = &PendingRequest{waitCh: waitCh}
	mu.Unlock()

	fmt.Printf("\ncalculating initial data for company: %s\n", companyId)
	time.Sleep(1 * time.Second)

	id, err := strconv.Atoi(companyId)
	if err != nil {
		fmt.Printf("invalid companyId: %v\n", err)
		return 0
	}

	data := id / 100

	mu.Lock()
	initialDataCache[companyId] = data
	delete(initialDataPending, companyId)
	mu.Unlock()

	waitCh <- data
	close(waitCh)
	return data
}
