package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// senior task
// Реализовать функцию processParallel
// Прокинуть контекст

type outVal struct {
	val int
	err error
}

var errTimeOut = errors.New("timed out")

func processData(ctx context.Context, val int) chan outVal {
	ch := make(chan struct{})
	out := make(chan outVal)

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()

	go func() {
		select {
		case <-ch:
			out <- outVal{
				val: val * 2,
				err: nil,
			}
		case <-ctx.Done():
			out <- outVal{
				val: 0,
				err: errTimeOut,
			}
		}
	}()

	return out
}

func main() {
	in := make(chan int)
	out := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		//time.Sleep(10 * time.Second)
		defer close(in)
		for i := range 10 {
			select {
			case in <- i + 1:
			case <-ctx.Done():
				return
			}
		}
	}()

	now := time.Now()
	processParallel(ctx, in, out, 5)

	for val := range out {
		fmt.Println("v =", val)
	}

	fmt.Println("main duration:", time.Since(now))
}

// Операция должна выполняться не более 5 секунд
func processParallel(ctx context.Context, in <-chan int, out chan<- int, numWorkers int) {
	wg := sync.WaitGroup{}
	for range numWorkers {
		wg.Add(1)
		go workerPool(ctx, in, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func workerPool(ctx context.Context, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case v, ok := <-in:
			if !ok {
				return
			}
			select {
			case ov := <-processData(ctx, v):
				if ov.err != nil {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- ov.val:
				}
			}

		case <-ctx.Done():
			return
		}
	}

}
