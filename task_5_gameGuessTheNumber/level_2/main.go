package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type level struct {
	levelName   string
	maxAttempts int
	randomNum   int
}

func main() {

	// Цыетовые политры
	winColor := color.New(color.FgGreen).SprintFunc()
	lossColor := color.New(color.FgRed).SprintFunc()
	attemptsColor := color.New(color.FgYellow).SprintFunc()

	fmt.Println()
	fmt.Println(winColor("Игра 'Угадай число'"))
	fmt.Println()

	// Уровни сложности
	levels := []level{
		{levelName: "Easy", maxAttempts: 15, randomNum: RandomNumberGenerator(50)},
		{levelName: "Medium", maxAttempts: 10, randomNum: RandomNumberGenerator(100)},
		{levelName: "Hard", maxAttempts: 5, randomNum: RandomNumberGenerator(200)},
	}

	attempts := 0
	systemReader := bufio.NewReader(os.Stdin) // подключаем ридера для чтения текста из терминала
	listPlayerNum := make([]int, 0)           // Срез для записи введенных чисел игрока
	var selectLevel *level

	fmt.Println("Уровни сложности:")
	fmt.Println(" * Easy: 1–50, 15 попыток")
	fmt.Println(" * Medium: 1–100, 10 попыток")
	fmt.Println(" * Hard: 1–200, 5 попыток")
	fmt.Println()
	fmt.Printf("Выбери уровень сложности: ")

	inputLevel, _ := systemReader.ReadString('\n')
	inputLevel = strings.TrimSpace(inputLevel)

	for index, _ := range levels {
		if inputLevel == levels[index].levelName {
			selectLevel = &levels[index]
			break
		}
	}

	for selectLevel.maxAttempts > attempts {
		fmt.Printf("Попытка #%s - Введите число: ", attemptsColor(attempts+1)) // счетчик попыток увеличиваем на 1

		input, _ := systemReader.ReadString('\n') // записываем все что написали до нажатия на enter
		input = strings.TrimSpace(input)          // редактируем получившийся результат, убираем пробелы и переносы
		playerNumber, err := strconv.Atoi(input)  // преобразовываем строку в число
		if err != nil {
			fmt.Printf("Ошибка: введена строка вместо числа. Введено: %s\n", input)
			return
		}

		listPlayerNum = append(listPlayerNum, playerNumber)
		fmt.Println("Твои числа:", attemptsColor(listPlayerNum))
		attempts++

		switch {
		case selectLevel.randomNum == playerNumber:
			fmt.Println()
			fmt.Println(winColor("Вы угадали число! 🎉 Игра закончена!"))
			fmt.Println()
			return
		case int(math.Abs(float64(selectLevel.randomNum-playerNumber))) <= 5:
			fmt.Println("Горячо 🔥")
			fmt.Println()
		case int(math.Abs(float64(selectLevel.randomNum-playerNumber))) <= 15:
			fmt.Println("Тепло 🙂")
			fmt.Println()
		default:
			fmt.Println("Холодно ❄️")
			fmt.Println()
		}

		if attempts == selectLevel.maxAttempts {
			fmt.Print(lossColor("Вы проиграли!😢 "))
			fmt.Println(lossColor("Загаданное число было:"), selectLevel.randomNum)
		}
	}
}

// Генератор рандомного числа.
func RandomNumberGenerator(numbers int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(numbers) + 1
	return randomNumber
}
