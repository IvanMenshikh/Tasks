package main

import (
	"fmt"
)

func sumZero(n int) ([]int, error) {
	if n < 1 || n > 1000 {
		return nil, fmt.Errorf("n должен быть в диапазоне от 1 до 1000, но получено: %d", n)
	}

	res := make([]int, 0, n)

	// Добавляем числа от 1 до n/2 и их противоположные значения
	for i := 1; i <= n/2; i++ {
		res = append(res, i)
		res = append(res, -i)
	}

	// Если n нечётное, добавляем 0
	if n%2 != 0 {
		res = append(res, 0)
	}

	return res, nil
}

func main() {
	n := 3
	result, err := sumZero(n)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(result) // пример вывода
}
