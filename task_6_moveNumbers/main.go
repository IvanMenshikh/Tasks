package main

import "fmt"

// двухуказательный алгоритм (two-pointer, stable compaction)
// Сложность: O(n) по времени, O(1) по памяти, без доп. аллокаций.

// вход [2, 4, 7, 0, 8, 9, 0]
// Вывод [2, 4, 7, 8, 9, 0, 0]

func moveNumbers(nums []int) []int {
	pos := 0 // индекс
	for _, num := range nums {
		if num != 0 {
			// Если число не ноль, перемещаем его на текущую позицию
			nums[pos] = num
			// Увеличиваем позицию для следующего ненулевого числа
			pos++
		}
	}
	// [2, 4, 7, 8, 9, 9, 0]
	fmt.Println("Промежуточный результат функции:", nums)

	// Добиваем последнии элементы нулями
	for pos < len(nums) {
		// 9 -> 0, т.к. pos = 5
		// 0 -> 0, т.к. pos = 6
		nums[pos] = 0
		pos++
	}

	return nums
}

func main() {

	slice := []int{2, 4, 7, 0, 8, 9, 0}
	fmt.Println("Вход:", slice)
	result := moveNumbers(slice)
	fmt.Println("Выход:", result)

}
