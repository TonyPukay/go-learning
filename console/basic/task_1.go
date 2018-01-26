package main

import (
	"fmt"
)

func main() {

	var num_1, num_2 int //input number
	var operation string //input operation

	var result int //result of operation
	var error bool //error of operation

	fmt.Print("Enter firts number:\t")
	fmt.Scanln(&num_1)

	fmt.Print("Enter second number:\t")
	fmt.Scanln(&num_2)

	fmt.Print("Enter operation:\t")
	fmt.Scanln(&operation)

	result, error = arithmetic(num_1, num_2, operation)

	if error {
		fmt.Println("ERROR: invalid operation!")
	} else {
		fmt.Println("Your result: ", result)
	}
}

func arithmetic(num_1, num_2 int, operation string ) (int, bool)  {

	var res int = 0
	var err bool = false

	switch operation {
	case "+":
		res = num_1 + num_2
	case "-":
		res = num_1 - num_2
	case "*":
		res = num_1 * num_2
	case "/":
		res = num_1 / num_2
	default:
		err = true
	}

	return res, err
}

