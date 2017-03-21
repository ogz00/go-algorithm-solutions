package main

import (
	"fmt"
	"bytes"
)

func main () {
	var min int
	if m, err := fmt.Scan(&min); m != 1 {
		panic(err)
	}

	if (min == 2) {
		fmt.Println("min(int, int)")
	}
	var buffer bytes.Buffer
	if (min >2) {
		for i := min; i>2 ; i-- {
			buffer.WriteString("min(int, ")
		}
		buffer.WriteString("min(int, int)")
		for i := min; i>2 ; i-- {
			buffer.WriteString(")")
		}
	}
	fmt.Println(buffer.String())

}
