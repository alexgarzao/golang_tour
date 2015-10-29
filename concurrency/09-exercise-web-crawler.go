package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(
	url string,
	depth int,
	fetcher Fetcher,
	crawled map[string]bool,
	ret chan string,
) {
	defer close(ret)
	if depth <= 0 {
		return
	}
	if crawled[url] {
		return
	}

	crawled[url] = true

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ret <- err.Error()
		return
	}

	ret <- fmt.Sprintf("found: %s %q", url, body)

	result := make([]chan string, len(urls))

	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, crawled, result[i])
	}

	for i := range result {
		for s := range result[i] {
			ret <- s
		}
	}

	return
}

func main() {
	crawled := make(map[string]bool)
	result := make(chan string)
	go Crawl("http://golang.org/", 4, fetcher, crawled, result)

	for s := range result {
		fmt.Println(s)
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
