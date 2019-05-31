package main

import (
	"bufio"
	//	"container/heap"
	"fmt"
	"os"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	text, _ := reader.ReadString('\n')
	return text
}

type Url struct {
	urlString string
	depth     int
}

type UrlQ []Url

func (q UrlQ) Len() int { return len(q) }

func (q UrlQ) Less(i, j int) bool {
	return q[i].depth < q[j].depth
}

func (q UrlQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].depth = i
	q[j].depth = j
}

func (q *UrlQ) Push(x interface{}) {
	n := len(*q)
	item := x.(*Url)
	item.depth = n
	*q = append(*q, item)
}

func (q *UrlQ) Pop() interface{} {
	return nil
}

type UrlFetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func crawl(start string, fetcher UrlFetcher) {
	urlHeap := make(UrlQ, 1, 100)
	urlHeap[0] = Url{start, 0}
	fmt.Println("Starting from URL:" + start)
	for len(urlHeap) > 0 {
		url := urlHeap[0]
		_, urls, _ := fetcher.Fetch(url.urlString)
		for i := 0; i < len(urls); i++ {
			// todo
			fmt.Println("next URLS:" + urls[i])
		}
		urlHeap.Pop()
	}
}

func main() {
	fmt.Println("Web crawler.")
	url := getInput("Gimme a URL:")
	fmt.Println("Got url:" + url)
}
