package main

import (
	"fmt"
	"sort"
)

func main() {
	var containers int
	if m, err := fmt.Scan(&containers); m != 1 || (m < 1 || m > 10) {
		panic(err)
	}
	for i := 0; i < containers; i++ {
		var ballTypes int
		if m, err := fmt.Scan(&ballTypes); m != 1 || (m < 1 || m > 10) {
			panic(err)
		}
		row := make([]int, ballTypes)
		col := make([]int, ballTypes)

		for i := 0; i < ballTypes; i++ {
			for j := 0; j < ballTypes; j++ {
				var x int
				if m, err := fmt.Scan(&x); m != 1 || (m < 1 || m > 100) {
					panic(err)
				}
				row[i] += x
				col[j] += x
			}
		}
		sort.Ints(row)
		sort.Ints(col)
		ok := true
		for i := 0; i < ballTypes; i++ {
			if (row[i] != col[i]) {
				ok = false
				break
			}
		}
		if (ok) {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}

	}

}
