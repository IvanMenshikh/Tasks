package main

import (
	"fmt"
	"sync"
)

// Конвейер чисел
// 1. Напиши программу, которая выполняет следующие шаги:
//	a. Функция GenerateNumbers
//	- Принимает целое число n.
//  - Генерирует числа от 1 до n и отправляет их в канал.
//
//	b. Функция SquareNumbers
//	- Принимает входной канал и выходной канал.
//	- Читает числа из входного канала, возводит их в квадрат и отправляет в выходной канал.
//
//	c. Функция PrintNumbers
//	- Принимает канал и печатает все числа, которые приходят в канал.
//
// 2. Соедини функции в конвейер: GenerateNumbers -> SquareNumbers -> PrintNumbers.
//
// 3. Используй горутины для генерации и обработки чисел, а каналы для передачи данных между ними.

func GenerateNumbers(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func SquareNumbers(in <-chan int, workers int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- v * v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func PrintNumbers(out <-chan int) {
	for v := range out {
		fmt.Println(v)
	}
}

func main() {

	nums := GenerateNumbers(100)
	squares := SquareNumbers(nums, 4)
	PrintNumbers(squares)

}
