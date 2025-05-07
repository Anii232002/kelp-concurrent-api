package calculations

import "sync"

type PendingRequest struct {
	waitCh chan int
}

var (
	mu                 sync.Mutex
	pendingRequests    = make(map[string]*PendingRequest)
	initialDataCache   = make(map[string]int)
	initialDataPending = make(map[string]*PendingRequest)
	apiResultCache     = make(map[string]int)
)
