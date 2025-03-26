package Tasks

import "fmt"

func TaskNine() {
	firstExercise()
	//Аналогично можно вызвать и другие функции
}

// 1
// Вынес в отдельное место
func firstExercise() {
	size := 5 // например, размер 5
	firstArray := make([]string, size)

	for i := 0; i < size; i++ {
		word := "x"
		firstArray[i] = word
		if i >= 1 {
			firstArray[i] = repeatSymbol(word, i)
		} else {
			firstArray[i] = word
		}
	}

	// Вывод результата
	fmt.Println(firstArray)
}

// Данная функция необходима для добавления +x символ с каждой новой итерацией
func repeatSymbol(str string, index int) string {
	wordFinally := ""
	//Конкатенация
	for i := 0; i <= index; i++ {
		wordFinally += str
	}
	return wordFinally
}

// 2
func arrayFill(str string, length int) {
	size := length
	array := make([]string, size)
	for i := 0; i < size; i++ {
		array[i] = str
	}

	fmt.Println(array)
}

// 3
func twoDimensionalArray() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5},
		{6},
	}
	sum := 0
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			sum += matrix[i][j]
		}
	}

	fmt.Println(sum)
}

// 4
func createArrayWithLoop() {
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3) // Инициализируем каждую строку размером 3
	}

	value := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = value
			value++
		}
	}

	fmt.Println(matrix)
}

// 5
func multiplicationOfArrayElements() {
	var array [4]int = [4]int{2, 5, 3, 9}
	result := (array[0] * array[1]) + (array[2] * array[3])
	fmt.Println(result)
}

// 6
//В Go для выполнения данной задачи подойдёт type

type Person struct {
	Name       string
	Surname    string
	Patronymic string
}

func personArray() {
	personOne := Person{
		Name:       "Kacey",
		Surname:    "Lee",
		Patronymic: "Musgraves",
	}
	personTwo := Person{
		Name:       "Taylor",
		Surname:    "Alison",
		Patronymic: "Swift",
	}

	persons := []Person{personOne, personTwo}
	fmt.Println(persons)
}

// 7
type Today struct {
	Year  string
	Month string
	Day   string
}

func todayArray() {
	dayToday := Today{
		Year:  "2025",
		Month: "Mart",
		Day:   "22",
	}
	fmt.Println(dayToday.Year)
	fmt.Println(dayToday.Month)
	fmt.Println(dayToday.Day)
}

// 8
func countElementsArray() {
	var alphabet []string = []string{"a", "b", "c", "d", "e", "f", "g"}
	var count int = 0
	for i := 0; i < len(alphabet); i++ {
		count += len(alphabet[i])
	}
	fmt.Println(count)
}

// 9
func endElementsArray() {
	var alphabet []string = []string{"a", "b", "c", "d", "e"}
	fmt.Println("Последний элемент массива: ", alphabet[len(alphabet)-1])
	fmt.Println("Предпоследний элемент массива: ", alphabet[len(alphabet)-2])
}
