package main

import "fmt"

// Shift сдвигает элементы слайса на n позиций вправо.
// Если n > len(slice), используется остаток от деления n % len(slice).
// Функция не изменяет исходный слайс.
func Shift(slice []int, n int) []int {

	result := make([]int, len(slice))

	n = n % len(slice) // нормализация сдвига

	copy(result[:n], slice[len(slice)-n:])
	fmt.Println("after copy 1", result)
	copy(result[n:], slice[:len(slice)-n])
	fmt.Println("after copy 2", result)
	return result
}

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("original slice", s)
	Shift(s, 2)
	fmt.Println("after shift", s)
}
