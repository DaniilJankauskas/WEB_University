package Tasks

import "fmt"

func TaskEleven() {
	ImageWithLoop()
}

func ImageWithLoop() {
	for i := 1; i <= 20; i++ {
		for j := 0; j < 20-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
