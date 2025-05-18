package main

import (
	"fmt"
)

// Функция для поиска максимального элемента
// Напишите функцию FindMax, которая принимает срез любого типа
// (например, int, float64, string) и возвращает максимальный элемент из этого среза.
// Используйте дженерики для реализации функции.

type Ordered interface {
	int | float64 | string
}

func FindMax[T Ordered](slice []T, accamulator func(a, b T) bool) T {

	max := slice[0]
	for _, val := range slice {
		if accamulator(val, max) {
			max = val
		}
	}

	return max
}

func main() {
	intSlice := []int{1, 3, 5, 2, 4}
	floatSlice := []float64{1.1, 3.3, 2.2, 5.5, 4.4}
	stringSlice := []string{"apple", "orange", "banana"}

	maxInt := FindMax(intSlice, func(a, b int) bool { return a > b })
	maxFloat := FindMax(floatSlice, func(a, b float64) bool { return a > b })
	maxString := FindMax(stringSlice, func(a, b string) bool { return a > b })

	fmt.Println("Max int:", maxInt)       // Max int: 5
	fmt.Println("Max float:", maxFloat)   // Max float: 5.5
	fmt.Println("Max string:", maxString) // Max string: orange
}
