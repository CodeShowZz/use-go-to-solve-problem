package main

import "fmt"

// 目的:go开启两个线程计算两个数组的和
func main() {
	c1 := make(chan int)
	go sum([]int{1, 2, 3}, c1)
	c2 := make(chan int)
	go sum([]int{4, 5, 6}, c2)
	sum := <-c1 + <-c2
	fmt.Println("两个线程分别进行计算,最终和为", sum)
}

func sum(a []int, ch chan int) {
	sum := 0
	for _, i := range a {
		sum += i
	}
	ch <- sum
}
