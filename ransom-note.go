package main

import "fmt"

func main() {
	var first int
	if m, err := fmt.Scan(&first); m != 1 {
		panic(err)
	}
	var second int
	if m, err := fmt.Scan(&second); m != 1 {
		panic(err)
	}

	fWords := createMaps(first)
	sWords := createMaps(second)
	ransom := false
	for k, v := range sWords.found {

		if _, ok := fWords.found[k]; !ok {
			ransom = false
			break
		} else if (v > fWords.found[k]) {
			ransom = false
			break
		} else {
			ransom = true
		}
	}
	if ransom {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

type words struct {
	found map[string]int
}

func newWords() *words {
	return &words{found:map[string]int{}}
}
func (w *words) add(word string, add int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = add
		return
	}
	w.found[word] = count + add
}

func createMaps(len int) *words {
	dict := newWords()
	for i := 0; i < len; i++ {
		var word string
		fmt.Scan(&word)
		dict.add(word, 1)
	}
	return dict
}