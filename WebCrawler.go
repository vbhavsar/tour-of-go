package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, 
		  body chan string,
		  urls chan []string,
		  error chan error) ()
}

var m map[string]bool

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	
	if m[url] {
		return
	}
	
	body := make(chan string)
	urls := make(chan []string)
	errChan := make(chan error)
	go fetcher.Fetch(url, body, urls, errChan)

	for {
		select {
		case bdy := <-body:
			fmt.Printf("found: %s %q\n", url, bdy)
			for _, u := range <-urls {
				Crawl(u, depth-1, fetcher)
			}
		case err := <- errChan:
			fmt.Println(err)
			return
		}
		break
	}
	
	return
}

func main() {
	m = make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, 
		  body chan string,
		  urls chan []string,
		  err chan error) () {
	m[url] = true
	if res, ok := f[url]; ok {
		body <- res.body
		urls <- res.urls
	} else {
		err <- fmt.Errorf("not found: %s", url)
	}
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

