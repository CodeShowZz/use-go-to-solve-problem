package main

import "fmt"

func compare(v, w int) bool {
	return v > w
}

func exchange(arrs []int, i int, j int) {
	temp := arrs[i]
	arrs[i] = arrs[j]
	arrs[j] = temp
}

func shellSort(arrs []int) {
	length := len(arrs)
	h := 1
	for h < length/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && !compare(arrs[j], arrs[j-h]); j -= h {
				exchange(arrs, j, j-1)
			}
		}
		h = h / 3
	}
}

func main() {
	arrs := []int{6, 5, 7, 2, 4, 1}
	shellSort(arrs)
	fmt.Println(arrs)
}
