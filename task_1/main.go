package main

import "fmt"

// Необходимо определить userIds участников, которые
// прошли наибольшее количество шагов (steps) за все дни,
// не пропустив ни одного дня соревнований.
// Пример
// # Пример 1

// # ввод:
// statistics = [
//     [ {userId: 1, steps: 1000}, {userId: 2, steps: 1500} ],
//     [ {userId: 2, steps: 1000} ]
// ]

// # вывод:
// champions = { userIds: [2], steps: 2500 }

// # Пример 2

// # ввод:
// statistics = [
//     [ {userId: 1, steps: 2000}, {userId: 2, steps: 1500} ],
//     [ {userId: 2, steps: 4000}, {userId: 1, steps: 3500} ]
// ]

// # вывод:
// champions = { userIds: [1, 2], steps: 5500 }

// Мое решение данной задачи:
type Record struct {
	UserId int
	Steps  int
}

type Result struct {
	UserIds []int
	Steps   int
}

func main() {

	statistics := [][]Record{
		{
			{UserId: 1, Steps: 2000},
			{UserId: 2, Steps: 4000},
		},
		{
			{UserId: 2, Steps: 1500},
			{UserId: 1, Steps: 3500}, // если убрать, получим решение #1
		},
	}

	//fmt.Println(statistics)

	var mapper = make(map[int]int)

	for _, day := range statistics {
		//fmt.Printf("Day: %d \n", i+1)
		for _, record := range day {
			if _, ok := mapper[record.UserId]; !ok {
				mapper[record.UserId] = 0
			}
			mapper[record.UserId] += record.Steps
		}
	}

	var result = Result{
		UserIds: []int{},
		Steps:   0,
	}

	//fmt.Println(mapper)
	for _, steps := range mapper {
		if steps > result.Steps {
			result.Steps = steps
		}
	}

	for id, steps := range mapper {
		if steps == result.Steps {
			result.UserIds = append(result.UserIds, id)
		}
	}
	fmt.Printf("champions = %+v\n", result)
}
