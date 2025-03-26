package Tasks

import (
	"fmt"
	"math"
	"math/rand"
)

func TaskSix() {
	//1
	var a, b int
	a = 10
	b = 3
	fmt.Println(a % b)

	//2
	if a%b == 0 {
		fmt.Println("Делится", a/b)
	} else {
		fmt.Println("Делится с остатком: ", a%b)
	}

	//3
	numberPow := 2.0
	st := math.Pow(numberPow, 10.0)

	//4
	fmt.Println("Число в 10-й степени: ", st)

	//5
	fmt.Println("Корень из 245 = ", math.Sqrt(245))

	//6
	var numbers [7]float64 = [7]float64{4, 2, 5, 19, 13, 0, 10}
	var sum float64
	for _, number := range numbers {
		sum += number * number
	}
	fmt.Println("Корень из суммы квадратов элементов массива: ", math.Sqrt(sum))

	//7
	floorNumber := math.Sqrt(379)
	fmt.Printf("Квадратный корень из 379: \nЦелое: %d\nДесятые: %.1f\nСотые: %.2f\n",
		int(math.Round(floorNumber)),
		math.Round(floorNumber*10)/10,
		math.Round(floorNumber*100)/100)

	//8
	var numbersMinMax [7]float64 = [7]float64{4, -2, 5, 19, -130, 0, 10}
	minNumber, maxNumber := math.MaxFloat64, -math.MaxFloat64
	for _, num := range numbersMinMax {
		if num < minNumber {
			minNumber = num
		}
		if num > maxNumber {
			maxNumber = num
		}
	}
	fmt.Printf("Минимальное: %.2f, Максимальное: %.2f\n", minNumber, maxNumber)

	//9
	fmt.Println("Случайное число от 0 до 100", rand.Intn(101))

	//10
	var numbersRandomArray [9]int
	for i, _ := range numbersRandomArray {
		numberGen := rand.Int()
		numbersRandomArray[i] = numberGen
	}
	for i, n := range numbersRandomArray {
		fmt.Println(i, n)
	}

	//11
	numbA := 15
	numbB := 900
	fmt.Println(math.Abs(float64(numbA) - float64(numbB)))

	//12
	var positiveNegativeNumber [6]int = [6]int{1, 2, -1, -2, 3, -3}
	fmt.Println("До: ")
	for i, number := range positiveNegativeNumber {
		fmt.Println(i, number)
	}
	for i, number := range positiveNegativeNumber {
		if number < 0 {
			positiveNegativeNumber[i] = positiveNegativeNumber[i] * (-1)
		}
	}
	fmt.Println("После: ")
	for i, number := range positiveNegativeNumber {
		fmt.Println(i, number)
	}

	//13
	numberDel := 30 //Можно через Scan делать, чтобы вводить числа
	fmt.Println("Делители числа: ")
	var numberDelNumbers []int
	for i := 1; i <= numberDel; i++ {
		if numberDel%i == 0 {
			numberDelNumbers = append(numberDelNumbers, i)
		}
	}
	fmt.Println(numberDelNumbers)

	//14
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumElementsArray := 0
	count := 0

	for _, value := range array {
		sumElementsArray += value
		count++

		if sumElementsArray > 10 {
			break
		}
	}
	fmt.Printf("Количество первых элементов, которые нужно сложить: %d\n", count)
}
