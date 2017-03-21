package main

import "fmt"

func main() {
	var first string
	if m, err := fmt.Scan(&first); m != 1 {
		panic(err)
	}

	var second string
	if m, err := fmt.Scan(&second); m != 1 {
		panic(err)
	}

	firstMap := make(map[rune]int, len(first))
	secondMap := make(map[rune]int, len(second))
	for _, v := range first {
		if _, ok := firstMap[v]; ok {
			firstMap[v] ++
		} else {
			firstMap[v] = 1
		}
	}
	for _, v := range second {
		if _, ok := secondMap[v]; ok {
			secondMap[v] ++
		} else {
			secondMap[v] = 1
		}
	}
	deleteCount := 0

	for k, v := range firstMap {
		if _, ok := secondMap[k]; !ok {
			deleteCount += v
		}else if(v != secondMap[k]){
			diff := v-secondMap[k]
			if(diff <0) {diff = -diff}
			deleteCount += diff
		}
	}
	for k, v := range secondMap {
		if _, ok := firstMap[k]; !ok {
			deleteCount += v
		}
	}
	fmt.Println(deleteCount)
}
