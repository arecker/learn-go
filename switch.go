package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	today := time.Now().Weekday()
	fmt.Println("Today is", today)
	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("go to work, shlub")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's the morning, technically")
	default:
		fmt.Println("It's the afternoon, technically")
	}

}
