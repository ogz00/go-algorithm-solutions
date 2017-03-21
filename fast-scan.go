package main

import "fmt"

func main() {
	fmt.Println(`Enter the integers`)
	all := make([]int, 3)
	ReadN(all, 0, 3)
	fmt.Println(all)
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}