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

func insertSort(arrs [] int) {
	length := len(arrs)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && compare(arrs[j-1], arrs[j]); j-- {
			exchange(arrs, j, j-1)
		}
	}
}

func main() {
	arrs := []int{6, 5, 7, 2, 4, 1}
	insertSort(arrs)
	fmt.Println(arrs)
}
