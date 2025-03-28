package Tasks

import (
	"fmt"
	"regexp"
)

func TaskOne() {
	controlWord := "'ahb acb aeb aeeb adcb axeb"

	//a + 2 символа + b
	fmt.Println(findCombinationSymbolsInWord(controlWord, 'a', 'b'))

	//a + 2 символа + a
	fmt.Println(findCombinationSymbolsInWord(controlWord, 'a', 'a'))
}

// Данная функция более универсальная
func findCombinationSymbolsInWord(controlWord string, charOne rune, charTwo rune) bool {
	pattern := fmt.Sprintf(".*%c.{2}%c.*", charOne, charTwo)
	condition := regexp.MustCompile(pattern)
	return condition.MatchString(controlWord)
}
