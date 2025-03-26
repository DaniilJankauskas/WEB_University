package Tasks

import "fmt"

func TaskFive() {
	myNum := 5
	answer := myNum
	fmt.Println("Проверяем присвоение к answer: ", answer)
	answer += 2
	fmt.Println("Проверяем answer +=2 : ", answer)
	answer *= 2
	fmt.Println("Проверяем answer *=2 : ", answer)
	answer -= 2
	fmt.Println("Проверяем answer -=2 : ", answer)
	answer /= 2
	fmt.Println("Проверяем answer /=2 : ", answer)
	fmt.Println("Разность исходного числа и answer: ", myNum-answer)
}
