package main

import "fmt"

// вход [2, 4, 7, 0, 8, 9, 0]
// Вывод [2, 4, 7, 8, 9, 0, 0]

func moveNumbers(nums []int) []int {
	pos := 0
	for _, num := range nums {
		if num != 0 {
			nums[pos] = num
			pos++
		}
	}

	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
	return nums
}

func main() {

	slice := []int{2, 4, 7, 0, 8, 9, 0}
	result := moveNumbers(slice)

	fmt.Println(result)

}
