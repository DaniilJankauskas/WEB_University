package Tasks

import "fmt"

func TaskSeven() {
	str := "Sveiki" // Пример строки
	myNumLen := printStringReturnNumber(str)
	fmt.Printf("Длина строки: %d\n", myNumLen)
}

func printStringReturnNumber(str string) int {
	fmt.Println(str)
	return len(str)
}
