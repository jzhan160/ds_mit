package crawl

import (
	"sync"
)

func ConcurrentCrawl(url string, fetcher Fetcher, state *fetchState) {
	state.mu.Lock()
	exist := state.fetched[url]
	state.fetched[url] = true
	state.mu.Unlock()

	if exist {
		return
	}

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	var done sync.WaitGroup
	for _, u := range urls {
		done.Add(1)
		go func(u string) {
			defer done.Done()
			ConcurrentCrawl(u, fetcher, state)
		}(u)
	}
	done.Wait()
}

type fetchState struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func MakeState() *fetchState {
	f := &fetchState{}
	f.fetched = make(map[string]bool)
	return f
}
