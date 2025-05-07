package calculations

import "fmt"

func checkCacheWithLog(key string, cache map[string]int, label string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	if data, exists := cache[key]; exists {
		fmt.Printf("\n%s already cached for key: %s\n", label, key)
		return data, true
	}

	return 0, false
}
