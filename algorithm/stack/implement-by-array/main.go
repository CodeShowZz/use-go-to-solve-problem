package main

import "fmt"

type ResizingArrayStack struct {
	array  [] string
	length int
}

func (stack *ResizingArrayStack) isEmpty() bool {
	return stack.length == 0
}

func (stack *ResizingArrayStack) size() int {
	return stack.length
}

func (stack *ResizingArrayStack) resize(newSize int) {
	newArray := make([] string, newSize)
	for i := 0; i < stack.length; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
}

func (stack *ResizingArrayStack) push(item string) {
	if stack.length == len(stack.array) {
		stack.resize(2 * len(stack.array))
	}
	stack.array[stack.length] = item
	stack.length++
}

func (stack *ResizingArrayStack) pop() string {
	item := stack.array[stack.length-1]
	stack.array[stack.length-1] = ""
	stack.length--
	if stack.length > 0 && stack.length == len(stack.array)/4 {
		stack.resize(len(stack.array) / 2)
	}
	return item
}

func main() {
	stack := ResizingArrayStack{[]string{""}, 0}
	stack.push("a")
	stack.push("b")
	stack.push("c")
	fmt.Println(stack.size())
	temp := stack.pop()
	fmt.Println(temp)
	fmt.Println(stack.size())
}
