package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Игра 'Угадай число' - от 1 до 100 началась!")
	fmt.Println("Угадайте число за 10 попыток!")
	fmt.Println()

	randomNum := RandomNumberGenerator(100) // Генерируем рандомное число
	//fmt.Println("Рандомное число:", randomNum)

	maxAttempts := 10 // максимум попыток
	attempts := 0
	systemReader := bufio.NewReader(os.Stdin) // подключаем ридера для чтения текста из терминала

	for maxAttempts > attempts {
		fmt.Printf("Попытка #%d - Введите число: ", attempts+1) // счетчик попыток увеличиваем на 1

		input, _ := systemReader.ReadString('\n') // записываем все что написали до нажатия на enter
		input = strings.TrimSpace(input)          // редактируем получившийся результат, убираем пробелы и переносы
		playerNumber, _ := strconv.Atoi(input)    // преобразовываем строку в число

		attempts++

		if randomNum == playerNumber {
			fmt.Println("Вы угадали число! 🎉")
			fmt.Println("Игра закончена!")
			break
		} else if randomNum > playerNumber {
			fmt.Println("Секретное число больше 👆")
		} else {
			fmt.Println("Секретное число меньше 👇")
		}

		if attempts == maxAttempts {
			fmt.Printf("Вы проиграли!😢")
			fmt.Printf("Загаданное число было: %d\n", randomNum)
		}
	}
}

// Генератор рандомного числа.
func RandomNumberGenerator(numbers int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(numbers) + 1
	return randomNumber
}
