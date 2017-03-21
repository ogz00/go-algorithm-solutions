package main

import (
	"fmt"
	"math"
)

func main() {
	testCase := 0
	fmt.Scan(&testCase)
	for i := 0; i < testCase; i++ {
		var min, max float64
		fmt.Scan(&min)
		fmt.Scan(&max)
		count := isPerfectSquare(min, max)
		fmt.Println(count)

	}
}

func isPerfectSquare(min float64, max float64) float64 {
	count := 0.0
	p1 := math.Sqrt(min)
	if (p1 * p1 != min) {
		p1 += 1
	}
	for i := p1; i * i <= max; i++ {
		count ++
	}

	return count

}