package main

import "fmt"

func main() {
	var candyNumb int
	if m, err := fmt.Scan(&candyNumb); m != 1 {
		panic(err)
	}
	var peopleNumb int
	if m, err := fmt.Scan(&peopleNumb); m != 1 {
		panic(err)
	}

	removeCandies := make([]int, peopleNumb)
	for i := 0; i < peopleNumb; i++ {
		var remove int;
		if m, err := fmt.Scan(&remove); m != 1 {
			panic(err)
		}
		removeCandies[i] = remove;
	}

	var addValue = 0;
	var tmpCandy = candyNumb;
	for k, m := range removeCandies {
		tmpCandy -= m;
		if (k != peopleNumb - 1) {
			if (tmpCandy < 5) {
				ref := (candyNumb - tmpCandy)
				addValue += ref
				tmpCandy += ref
			}
		}
	}
	fmt.Printf("%d", addValue)
}
