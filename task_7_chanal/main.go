package main

import (
	"fmt"
	"time"
)

// Написать 3 функции:
// 1. Writer - генерит числа от 1 до 10
// 2. Doubler - умножает числа на 2 имитируя работу (500мс)
// 3. Reader - читает и выводит на экран

func Writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func Doubler(inputCh <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range inputCh {
			time.Sleep(500 * time.Millisecond) // иммитируем долгую работу
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func Reader(ch <-chan int) {
	for v := range ch {
		fmt.Println("Read Value:", v)
	}
}

func main() {
	Reader(Doubler(Writer()))
}
