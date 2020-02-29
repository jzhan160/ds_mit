package crawl


func SerialCrawl(url string, fetcher Fetcher, fetched map[string]bool) {
	if fetched[url] {
		return
	}
	fetched[url] = true
	_,urls,err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	for _, u := range urls {
		SerialCrawl(u, fetcher, fetched)
	}
	return
}