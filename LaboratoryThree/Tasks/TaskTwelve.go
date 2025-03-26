package Tasks

import (
	"fmt"
	"strconv"
)

func TaskTwelve() {
	middleVarCalculate() //К примеру
}

// 1
func middleVarCalculate() {
	array := []int{1, 2, 3, 4, 5}
	count := len(array)
	middleVar := countSum(array, count) / count
	fmt.Println(middleVar)
}

func countSum(array []int, count int) int {
	if count == 0 {
		return 0
	}
	var sum int
	sum += array[count-1] + countSum(array, count-1)
	fmt.Println(array[count-1])
	return sum
}

// 2
func sumNumbers() {
	var sum int = 0
	sigma := numberSeries(1, sum)
	fmt.Println(sigma)
}

func numberSeries(current int, sum int) int {
	if current > 99 {
		return sum
	}
	sum += current
	return numberSeries(current+1, sum)
}

// 3
func createAlphabetMap(letter rune, number int) map[string]int {
	if letter > 'z' {
		return map[string]int{}
	}
	alphabet := createAlphabetMap(letter+1, number+1)
	alphabet[string(letter)] = number
	return alphabet
}

// 4
func sumPairs(s string, index int) int {
	if index >= len(s) {
		return 0
	}
	pair := s[index : index+2]
	num, _ := strconv.Atoi(pair)
	return num + sumPairs(s, index+2)
}
