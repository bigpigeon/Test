package main

import (
	"fmt"
	"sync"
)

var cache map[string]bool

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func GatherUrls(url string, fetcher Fetcher, Urls chan []string) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("found: %s %q\n", url, body)
	}
	Urls <- urls
	wg.Done()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// get all urls for depth
	// check if url has been crawled
	//  Y: noop
	//  N: crawl url
	// when depth is 0, stop
	fmt.Printf("crawling %q... %d\n", url, depth)
	if depth <= 0 {
		return
	}
	uc := make(chan []string)
	wg.Add(1)
	go GatherUrls(url, fetcher, uc)
	urls, _ := <-uc
	fmt.Println("urls:", urls)
	for _, u := range urls {
		fmt.Println("currentUrl:", u)
		if _, exists := cache[u]; !exists {
			fmt.Printf("about to crawl %q\n", u)
			wg.Add(1)
			go Crawl(u, depth-1, fetcher)
		} else {
			cache[u] = true
		}
	}
	wg.Done()
}

func main() {
	cache = make(map[string]bool)
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
