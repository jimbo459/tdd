package concurrency

import (
	"fmt"
	"time"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
			fmt.Println("finished adding item")
		}(url)
	}

	time.Sleep(time.Second * 5)
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		fmt.Println(result)
		results[result.string] = result.bool
	}

	return results
}
