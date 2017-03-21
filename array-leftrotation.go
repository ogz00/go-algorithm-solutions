package main

import "fmt"

func main() {
	var arraySize int
	var rotation int
	if m, err := fmt.Scan(&arraySize); m != 1 || (m < 1 || m > 100000) {
		panic(err)
	}
	if m, err := fmt.Scan(&rotation); m != 1 || (m < 1 || m > arraySize) {
		panic(err)
	}

	myArray := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		if m, err := fmt.Scan(&myArray[i]); m != 1 || (m < 1 || m > 1000000) {
			panic(err)
		}
	}

	for _,v := range rotate(myArray, rotation) {
		fmt.Printf("%d ", v)
	}

}

func rotate(arr []int, rot int) [] int{
	newArr := make([]int, len(arr))
	for k, v := range arr {
		indice := k - rot
		if ( indice >= 0) {
			newArr[indice] = v
		}else{
			newArr[len(arr) + indice] = v
		}
	}

	return newArr

}
