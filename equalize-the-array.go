package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   int
	Value int
}
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	var arrLen int
	if m, err := fmt.Scan(&arrLen); m != 1 || (m < 1 || m > 100) {
		panic(err)
	}

	uerSlice := make([] int, arrLen)
	for i := 0; i < arrLen; i++ {
		if m, err := fmt.Scan(&uerSlice[i]); m != 1 || (m < 1 || m > 100) {
			panic(err)
		}
	}
	fmt.Println(calcDeleteCount(findRepeat(uerSlice)))

}

func findRepeat(arr []int) map[int]int {
	repeatMap := make(map[int]int)
	for _, v := range arr {
		repeatMap[v]++
	}
	return repeatMap
}

func rankByIntCount(intFrequencies map[int]int) PairList {
	pl := make(PairList, len(intFrequencies))
	i := 0
	for k, v := range intFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func calcDeleteCount(repeatMap map[int]int) int {
	deleteCount := 0
	for k, p := range rankByIntCount(repeatMap) {
		if (k != 0) {
			deleteCount += p.Value
		}
	}

	return deleteCount
}



