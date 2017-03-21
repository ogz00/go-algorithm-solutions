package main

import ("fmt")

func main() {
	all := make(map[int]int, 26)
	for k :=97; k<123; k++{
		var inp int
		if m, err :=  fmt.Scan(&inp); m != 1 {
			panic(err)
		}
		all[k] = inp
	}

	var word string
	fmt.Scan(&word)
	wide := len(word)
	tallest := 1

	for _,v :=range word {
		if all[int(v)]> tallest {
			tallest = all[int(v)]
		}
	}

	area := wide *tallest
	fmt.Printf("%d", area);


}
