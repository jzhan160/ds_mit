package main

import (
	"fmt"
	"haha/DS/ThreadsAndRpc/crawl"
	"haha/DS/ThreadsAndRpc/rpc"
	"time"
)

func main() {
	var fetcher = crawl.FakeFetcher{
		"https://golang.org/": &crawl.FakeResult{
			Body: "The Go Programming Language",
			Urls: []string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &crawl.FakeResult{
			Body: "Packages",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/os/": &crawl.FakeResult{
			Body: "Package os",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
	startUrl := "https://golang.org/"

	fmt.Println("============")
	fmt.Println("serial crawl")
	fmt.Println("============")
	start1:= time.Now().Nanosecond()
	crawl.SerialCrawl(startUrl, fetcher,make(map[string]bool))
	cost1 := time.Now().Nanosecond()-start1
	fmt.Println("cost of serial crawl: ",cost1)

	fmt.Println("================")
	fmt.Println("concurrent crawl")
	fmt.Println("================")
	start2:= time.Now().Nanosecond()
	state := crawl.MakeState()
	crawl.ConcurrentCrawl(startUrl,fetcher,state)
	cost2 := time.Now().Nanosecond()-start2
	fmt.Println("cost of concurrent crawl: ",cost2)

	fmt.Println("=============================")
	fmt.Println("concurrent crawl with channel")
	fmt.Println("=============================")
	start3 := time.Now().Nanosecond()
	crawl.ConcurrentCrawlWithChannel(startUrl,fetcher)
	cost3 := time.Now().Nanosecond()-start3
	fmt.Println("cost of concurrent crawl with channel: ",cost3)


	rpc.Server()

	rpc.Put("subject", "6.824")
	fmt.Printf("Put(subject, 6.824) done\n")
	fmt.Printf("get(subject) -> %s\n", rpc.Get("subject"))
}
