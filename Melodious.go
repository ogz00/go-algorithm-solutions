package main

import (
	"fmt"
)


var wordCount = 0
var wordLen int
var pwd [] int

func main() {

	if m, err := fmt.Scan(&wordLen); m != 1 || (m < 1 || m > 6) {
		panic(err)
	}
	pwd = make([]int, wordLen)

	printInit(pwd)

}

func printInit(pwd []int) {
	for k := 97; k < 123; k++ {
		wordCount = 0
		if (k != 121) {
			pwd[wordCount] = k
			printRest(wordLen - 1)

		}

	}
}

func printRest(num int) {
	if num == 0 {
		printArray()
	} else {
		wordCount++
		if (isVowel(pwd[wordCount - 1])) {
			temp1 := wordCount

			for k := 97; k < 123; k++ {
				if (k != 121 && !isVowel(k)) {
					pwd[wordCount] = k
					printRest(num - 1)
				}
				wordCount = temp1
			}
		} else {
			temp2 := wordCount
			vowels := []rune{'a', 'e', 'i', 'o', 'u'}

			for _,ch := range vowels {
				pwd[wordCount] = int(ch)
				printRest(num - 1)
				wordCount = temp2
			}
		}
	}
}

func printArray() {
	for i := 0; i < wordLen; i++ {
		fmt.Printf("%s", string(pwd[i]))
	}
	fmt.Println()
}

func isVowel(ch int) bool {
	if (rune(ch) == 'a' || rune(ch) == 'e' || rune(ch) == 'i' || rune(ch) == 'o' || rune(ch) == 'u') {
		return true
	} else {
		return false
	}

}
