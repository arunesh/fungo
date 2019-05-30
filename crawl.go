package main

import  (
        "fmt"
        "bufio"
	"container/heap"
	)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	text, _ := reader.ReadString('\n')
	return text
}

type Url struct {
      urlString string
      depth int
}

type UrlQ []Url

func (q UrlQ) Len() { return len(q) }

func (q UrlQ) Less(i, j int) bool {
     return q[i].depth < q[j].depth
}

func (q UrlQ) Swap(i, j int) {
}

func (q *UrlQ) Push(x interface{}) {
}

func (q *UrlQ) Pop() interface{} {
   return nil
}

type interface UrlFetcher {
    Fetch(url string) (body string, urls []string, err error)
}

func crawl(start string, fetcher UrlFetcher) {
     urlHeap := make(UrlQ, 1, 100)
     urlHeap[0] := Url{start, 0}
     fmt.Println("Starting from URL:" + start)
     while len(urlHeap) > 0 {
         url := urlHeap[0]
         body, urls, err := fetcher.Fetch(url.urlString)
	 for i = 0; i < len(urls); i ++ {
	    // todo
	 }
     }
}

func main() {
	fmt.Println("Web crawler.")
	url := getInput("Gimme a URL:")
	fmt.Println("Got url:" + url)
}
