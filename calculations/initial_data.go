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
		fmt.Println(companyId, " ", data)
		return data
	}

	return startInitialDataCalculation(companyId)

}

func startInitialDataCalculation(companyId string) int {

	mu.Lock()
	initialDataPending[companyId] = &PendingRequest{waitCh: []chan int{}}
	mu.Unlock()

	fmt.Printf("\ncalculating initial data for company: %s\n", companyId)
	time.Sleep(1 * time.Second)

	id, err := strconv.Atoi(companyId)
	if err != nil {
		fmt.Printf("invalid company id: %v\n", err)
		return 0
	}

	data := id / 100

	mu.Lock()
	initialDataCache[companyId] = data
	waitChs := initialDataPending[companyId].waitCh
	delete(initialDataPending, companyId)
	mu.Unlock()

	for _, ch := range waitChs {
		ch <- data
		close(ch)
	}

	return data
}
