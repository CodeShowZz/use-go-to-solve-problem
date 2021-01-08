package main

import "fmt"

func compare(v, w int) bool {
	return v > w
}

func exchange(arrs [] int, i int, j int) {
	temp := arrs[i]
	arrs[i] = arrs[j]
	arrs[j] = temp
}

func selectSort(arrs [] int) {
	length := len(arrs)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if compare(arrs[min], arrs[j]) {
				min = j
			}
		}
		exchange(arrs,i,min)
	}
}

func main() {
	arrs := []int{6,5,7,2,4,1}
	selectSort(arrs)
	fmt.Println(arrs)
}
