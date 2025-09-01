package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.example.com/p/mp002xw0lvkd/",
	"https://www.ya.ru/p/mp002xw14uf2/",
	"https://www.example.ru/p/rtladr746901/",
	"https://www.example.ru/p/mp002xw18h9d/",
	"https://www.example.ru/p/mp002xw004x4/",
	"https://www.example.ru/p/mp002xw0zfxy/",
}

type job struct {
	index int
	url   string
}

func crawl(urls []string, k int) []int {

	jobs := make(chan job)
	result := make([]int, len(urls))
	var wg sync.WaitGroup

	for i := 0; i < k; i++ {

		go func() {
			for j := range jobs {
				resp, err := http.Get(j.url)
				if err != nil {
					fmt.Println("Error:", err)
					result[j.index] = 0
				} else {
					result[j.index] = resp.StatusCode
					resp.Body.Close()
				}
				wg.Done()
			}
		}()

		for i, url := range urls {
			wg.Add(1)
			jobs <- job{
				index: i,
				url:   url,
			}
		}

		go func() {
			wg.Wait()
			close(jobs)
		}()
	}
	return result
}

func main() {
	result := crawl(urls, 5)
	fmt.Println("All done", result)
}
