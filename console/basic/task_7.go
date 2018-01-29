package main

import (
	"fmt"
)

func main() {
	var month, day uint8
	var year uint16
	var result bool

	fmt.Print("Enter day:\t")
	fmt.Scanln(&day)

	fmt.Print("Enter month:\t")
	fmt.Scanln(&month)

	fmt.Print("Enter year:\t")
	fmt.Scanln(&year)

	result = date(day, month, year)

	if result {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}

func date(day, month uint8, year uint16) bool {
	switch month {
	case 1,3,5,7,8,10,12:{
		if day >= 1 && day <= 31 {
			return true
		} else {
			return false
		}
	}
	case 4,6,9,11:{
		if day >= 1 && day <= 30 {
			return true
		} else {
			return false
		}
	}
	case 2:{
		if is_year_leap(int(year)) && day >= 1 && day <= 29 {
			return true
		} else if !is_year_leap(int(year)) && day >= 1 && day <= 28{
			return true
		} else {
			return false
		}
	}
	default:
		return false
	}
}

func is_year_leap(year int) bool {
	return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}
