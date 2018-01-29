package main

import (
	"fmt"
	"math"
)

func main() {

	var number int16
	var result bool

	fmt.Print("Enter your number: ")
	fmt.Scanln(&number)

	result = is_prime(number)

	if result {
		fmt.Print(number, " is prime")
	} else {
		fmt.Print(number, " is not prime")
	}
}

func is_prime(number int16) bool {

	for i := 2; int16(i) <= int16(math.Sqrt(float64(number))); i++ {
		if number%2 == 0{
			return false
		}
	}
	return true
}
