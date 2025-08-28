package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Имеется функция, которая работает неопределенно долго (до 100 секунд)
func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

// Написать обертку для этой функции, которая будет прерывать выполнение, если
// функция работает дольше 3 секунд, и возвращать ошибку.
func predictableTimeWork() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	timeStart := time.Now()
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
	case <-ctx.Done():
		fmt.Println("Вышли")
		fmt.Printf("Время работы: %v\n", time.Since(timeStart))
		return ctx.Err()
		// Вариант без контекста. Более компактное и простое решение
		// case <- time.After(3 * time.Second):
	}

	return nil
}

func main() {
	predictableTimeWork()
}
