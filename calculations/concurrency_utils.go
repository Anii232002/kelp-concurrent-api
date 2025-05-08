package calculations

import "fmt"

func waitForConcurrentRequest(key string, pendingList map[string]*PendingRequest, label string) (int, bool) {
	mu.Lock()
	pending, exists := pendingList[key]
	mu.Unlock()

	if exists {
		ch := make(chan int, 1)
		mu.Lock()
		pending.waitCh = append(pending.waitCh, ch)
		mu.Unlock()
		fmt.Printf("\nwaiting for other process to calculate %s for %s\n", label, key)
		return <-ch, true
	}
	return 0, false
}
