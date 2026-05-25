package scraper

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Result struct {
	URL   string
	Title string
	Err   error
}

func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title := extractTitle(c); title != "" {
			return title
		}
	}
	return ""
}

func fetchTitle(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	title := extractTitle(doc)
	ch <- Result{URL: url, Title: title}
}

func Scraper() {
	urls := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
		"https://news.ycombinator.com",
		"https://amazon.com",
		"https://linkedin.com",
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go fetchTitle(url, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	timeout := time.After(6 * time.Second)

	for {
		select {
		case res, ok := <-results:
			if !ok {
				return
			}
			if res.Err != nil {
				fmt.Printf("Error fetching %s: %v\n", res.URL, res.Err)
			} else {
				fmt.Printf("%s -> %s\n", res.URL, res.Title)
			}
		case <- timeout:
			fmt.Println("Timeout reached, exiting...")
			return
		}
	}

}
