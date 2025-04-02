package tasks

import (
	"fmt"
	"regexp"
	"strconv"
)

func TaskTwo() {
	text := "a1b2c3"
	multiplyNumberInWord(text)
}

func multiplyNumberInWord(controlWord string) {
	re := regexp.MustCompile("[0-9]+") //Как альтернатива \d
	matches := re.FindAllString(controlWord, -1)

	result := controlWord

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Ошибка преобразования", err)
			continue
		}
		cube := num * num * num
		result = regexp.MustCompile(match).ReplaceAllString(result, strconv.Itoa(cube))
	}
	fmt.Println(result)
}
