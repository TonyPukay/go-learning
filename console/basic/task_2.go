package main

import "fmt"

func main() {
	var year int
	var result bool

	fmt.Print("Enter year:\t")
	fmt.Scanln(&year)

	result = is_year_leap(year)

	if result {
		fmt.Print(year, " is leap")
	} else {
		fmt.Print(year, " is not leap")
	}
}

func is_year_leap(year int) bool {
	return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}
