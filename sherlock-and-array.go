package main

import "fmt"

func main() {
	arrNumb := 0
	fmt.Scan(&arrNumb)

	for i := 0; i < arrNumb; i++ {
		arrLen := 0
		fmt.Scan(&arrLen)
		myArr := make([]int, arrLen)
		for k := range myArr {
			fmt.Scan(&myArr[k])
		}
		if (sherlock(myArr)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func sherlock(arr []int) bool {
	if (len(arr) == 1) {
		return true
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	left := 0
	sum -= arr[0]
	for i := 1; i < len(arr); i++ {
		left += arr[i - 1]
		sum -= arr[i]
		if (left == sum) {
			return true
		}
	}
	return false
}