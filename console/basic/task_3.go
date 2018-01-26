package main

import "fmt"

func main() {

	var num_month int8
	var name_season, name_month string

	fmt.Print("Enter number of month: \t")
	fmt.Scanln(&num_month)

	name_season, name_month = season(num_month)

	fmt.Print(name_month, " is ", name_season)
}

func season(month int8) (string, string) {

	var name_season, name_month string

	if month >= 1 && month <= 2 || month == 12 {
		name_season = "winter"
	} else if month >= 3 && month <= 5{
		name_season = "spring"
	} else if month >= 6 && month <= 8{
		name_season = "summer"
	} else if month >= 9 && month <= 11 {
		name_season = "autumn"
	} else {
		name_season = "dich"
	}

	switch month {
	case 1: name_month = "January"
	case 2: name_month = "February"
	case 3: name_month = "March"
	case 4: name_month = "April"
	case 5: name_month = "May"
	case 6: name_month = "June"
	case 7: name_month = "July"
	case 8: name_month = "August"
	case 9: name_month = "September"
	case 10: name_month = "October"
	case 11: name_month = "November"
	case 12: name_month = "December"
	default:
		name_month = "dich"
	}



	return name_season, name_month
}
