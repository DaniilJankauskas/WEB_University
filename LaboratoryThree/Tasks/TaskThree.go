package Tasks

import "fmt"

func TaskThree() {
	var numLanguages, months, days int
	numLanguages = 4
	months = 11
	days = months * 16
	var daysPerLanguage float64 = float64(days) / float64(numLanguages)
	fmt.Println("Дни на языки: ", daysPerLanguage)
}
