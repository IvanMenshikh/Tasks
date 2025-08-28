package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wc sync.WaitGroup
	m := make(chan string, 3)
	for i := 0; i < 5; i++ {
		wc.Add(1)
		go func(mm chan<- string, i int) {
			defer wc.Done()
			mm <- fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
		}(m, i)
	}

	go func() {
		wc.Wait()
		close(m)
	}()

	for msg := range m {
		fmt.Println(msg)
	}

}
