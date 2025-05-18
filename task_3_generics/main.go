package main

import (
	"fmt"
)

// Функция для объединения двух срезов
// Напишите функцию JoinSlices, которая принимает два среза
// одного типа и возвращает новый срез, содержащий элементы обоих срезов.
// Используйте дженерики для реализации функции.

func JoinSlices[T any](list1, list2 []T) []T {
	newSlice := make([]T, 0, len(list1)+len(list2))
	newSlice = append(newSlice, list1...)
	newSlice = append(newSlice, list2...)
	return newSlice
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	fmt.Println(JoinSlices(slice1, slice2))
}
