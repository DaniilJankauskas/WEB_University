package Tasks

import "fmt"

func TaskEight() {
	//Пример вызова нескольких функций
	increaseEnthusiasm("Mundo")
	repeatIncreaseEnthusiasm("Mundo")
	fmt.Println(increaseEnthusiasm(repeatIncreaseEnthusiasm("Mundo")))

}

func increaseEnthusiasm(str string) string {
	//Плюс восклицательный символ
	return str + "" + "!"
}

func repeatIncreaseEnthusiasm(str string) string {
	//Конкатенация
	return str + str + str
}

func cut(s string, length ...int) string {
	maxLength := 10 //Значение по умолчанию =10
	if len(length) > 0 {
		maxLength = length[0]
	}
	// Если длина строки меньше данного значения, то просто возвращаем строку
	if len(s) <= maxLength {
		return s
	}

	// Иначе -> Возвращаем строку, обрезанную до 10
	return s[:maxLength]
}

// Рекурсивный обход массива
func recursionArray(arr []int, index int) {
	if index == len(arr) {
		return
	}
	fmt.Println(arr[index])
	recursionArray(arr, index+1)
}

// Функция для вычисления суммы цифр числа
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Функция для получения однозначной суммы цифр
func singleDigitSum(n int) int {
	for n > 9 {
		n = sumOfDigits(n)
	}
	return n
}
