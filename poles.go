package main

import (
	"fmt"
	"sort"
)

type pole struct {
	altitude int
	weight   int
}

type Cost struct {
	Key   int
	Value int
}
type CostList []Cost

func (p CostList) Len() int {
	return len(p)
}
func (p CostList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}
func (p CostList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	totalCost := 0
	var poles int
	if m, err := fmt.Scan(&poles); m != 1 || (m < 1 || m > 5000) {
		panic(err)
	}

	var stacks int
	if m, err := fmt.Scan(&stacks); m != 1 || (m < 1 || m > poles) {
		panic(err)
	}

	poleMap := make(map[int]pole, poles)
	costMap := make(CostList, poles)
	for i := 0; i < poles; i++ {
		p1 := pole{}
		if m, err := fmt.Scan(&p1.altitude); m != 1 || (m < 1 || m > 1000000) {
			panic(err)
		}
		if m, err := fmt.Scan(&p1.weight); m != 1 || (m < 1 || m > 1000000) {
			panic(err)
		}
		poleMap[i] = p1
		if (i != 0) {
			cost := (poleMap[i].altitude - poleMap[i - 1].altitude) * (poleMap[i].weight)
			costMap[i] = Cost{i, cost}
		}
	}

	if (stacks == 1) {
		for _, v := range costMap {
			totalCost += v.Key * v.Value
		}
		fmt.Println(totalCost)
	} else {
		fmt.Println(findTheBreaks(stacks, costMap))
	}

}

func findTheBreaks(stacks int, costList CostList) []int {
	sort.Sort(sort.Reverse(costList))
	fmt.Println(costList)
	breaks := make([]int, stacks-1)
	for i := 0; i < stacks-1; i++ {
		breaks[i] = costList[i].Key
	}
	return breaks
}
