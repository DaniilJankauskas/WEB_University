package Tasks

import "fmt"

func TaskOne() {
	veryBadUnclearName := "15 chicken wings"
	order := &veryBadUnclearName
	newOrder := *order + " " + " from Wendy's "
	fmt.Println(newOrder)
}
