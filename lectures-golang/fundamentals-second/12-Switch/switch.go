package main

import (
	"fmt"
)

func dayOfTheWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

// Another way to assess the condition of the Switch
// fallthrough: Transfers control to the next case, even if the current case may have matched.
// In Golang it is not necessary to use the Break clause
func dayOfTheWeekTwo(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func main() {
	day := dayOfTheWeek(200)
	fmt.Println(day)

	dayTwo := dayOfTheWeekTwo(5)
	fmt.Println(dayTwo)
}
