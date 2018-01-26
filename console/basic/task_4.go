package main

import (
	"fmt"
	"math"
)

func main() {
	var side float32
	parameter := make(map[string]float32)

	fmt.Print("Enter quare side: ")
	fmt.Scanln(&side)

	parameter["perimeter"], parameter["area"], parameter["diagonal"] = square(side)

	for item, value := range parameter {
		fmt.Println(item, "-", value)
	}

}

func square(side float32) (p float32, s float32, d float32) {
	var perimeter, area, diagonal float64

	perimeter = float64(side) * 4
	area = math.Pow(float64(side), 2)
	diagonal = math.Sqrt(float64(side) * 2)

	return float32(perimeter), float32(area), float32(diagonal)
}
