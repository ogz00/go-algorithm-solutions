package main

import ("fmt")

func main() {
	var count int
	var num []int
	var err error
	for _, err = fmt.Scanf("%d\n", &count); err == nil && count > 0; count-- {
		var input int
		_, err = fmt.Scanf("%d\n", &input)
		num = append(num, input)
	}
	for k, v := range num {
		if v % 5 > 2 && v > 37 {
			v += 5 - (v % 5)
			num[k] = v
		}
	}
	for _, v := range num {
		fmt.Printf("%d \n",v)
	}

}