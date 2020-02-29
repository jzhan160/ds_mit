package crawl

import "fmt"

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*FakeResult

type FakeResult struct {
	Body string
	Urls []string
}

func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		fmt.Printf("found:   %s\n", url)
		return res.Body, res.Urls, nil
	}
	fmt.Printf("missing: %s\n", url)
	return "", nil, fmt.Errorf("not found: %s", url)
}
