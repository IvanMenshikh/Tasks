package main

import (
	"fmt"
)

// Требуется написать функцию Filter, которая:
// Принимает слайс элементов любого типа ([]T) и функцию(func(T) bool)
// Возвращает новый слайс, содержащий только элементы, для которых функция вернула true
// Исходный слайс не должен изменяться и используйте дженерики для реализации функции

func Filter[T any](list []T, accumulator func(x T) bool) []T {

	newList := make([]T, 0, len(list))

	for _, v := range list {
		if accumulator(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	filtered := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println(filtered) // [2, 4]

	strList := []string{"banana", "orange", "grape", "watermelon"}
	filteredString := Filter(strList, func(x string) bool { return x == "grape" })
	fmt.Println(filteredString)
}
