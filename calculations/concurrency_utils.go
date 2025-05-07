package calculations

import "fmt"

func waitForConcurrentRequest(key string, pendingList map[string]*PendingRequest, label string) (int, bool) {
	mu.Lock()
	pending, exists := pendingList[key]
	mu.Unlock()

	if exists {
		fmt.Printf("\nwaiting for other process to calculate %s for %s\n", label, key)
		return <-pending.waitCh, true
	}
	return 0, false
}
