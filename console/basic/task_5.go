package main

import "fmt"

func main() {

	var amount, interest, result float32
	var year int16

	fmt.Print("Enter amount of money($ or ₴):\t")
	fmt.Scanln(&amount)

	fmt.Print("Enter percentage increase(%):\t")
	fmt.Scanln(&interest)

	fmt.Print("Enter period(year):\t")
	fmt.Scanln(&year)

	result, _, _ = bank(amount, interest, year)

	fmt.Print("Your result: ", result)
}

func bank(amount, interest float32, year int16) (float32, float32, int16) {

	if year == 0 {
		return amount, interest, year
	} else {
		//profit := amount * (interest/100)
		//new_amount := amount + profit
		new_amount := amount * (interest/100 + 1)
		year--
		return bank(float32(new_amount), interest, year)
	}
}
