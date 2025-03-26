package Tasks

import (
	"fmt"
	"strconv"
)

func TaskTen() {
	threeNumbersInArray()
}

// 1
func twoNumbersLargeTens(a int, b int) bool {
	c := a + b
	if c > 10 {
		return true
	} else {
		return false
	}
}

// 2
func twoNumbersEqual(a, b int) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

// 3
// К сожалению, в Go отсутствует поддержка тернарного оператора
func sorryTernarOperator() {
	word := 0
	if word == 0 {
		fmt.Println("Верно")
	}
}

// 4
func ageConditions() {
	age := 38
	if age <= 10 || age >= 99 {
		fmt.Println("Число в диапазоне меньше 10 или больше 99")
	} else if age > 10 || age < 99 {
		ageStr := strconv.Itoa(age)
		sumSymbol := 0

		for _, symbol := range ageStr {
			digit, _ := strconv.Atoi(string(symbol))
			sumSymbol += digit
		}

		fmt.Println("Возраст: ", age, "Сумма символов: ", sumSymbol)

		if sumSymbol <= 9 {
			fmt.Println("Сумма однозначна")
		} else {
			fmt.Println("Сумма двузначная")
		}
	}
}

// 5
func threeNumbersInArray() {
	array := []int{1, 2, 3, 4, 5}
	count := 0
	for i := 0; i < len(array); i++ {
		count++
	}
	if count >= 3 {
		sum := 0
		for i := 0; i < len(array); i++ {
			sum += array[i]
		}
		fmt.Println("Сумма элементов в массиве: ", sum)
	} else {
		fmt.Println("В массиве меньше трёх элементов")
	}
}
