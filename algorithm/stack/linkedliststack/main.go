package main

import "fmt"

type LinkedListStack struct {
	length int
	first  *Node
}

type Node struct {
	item string
	next *Node
}

func (stack *LinkedListStack) isEmpty() bool {
	return stack.length == 0
}

func (stack *LinkedListStack) size() int {
	return stack.length
}

func (stack *LinkedListStack) push(item string) {
	oldFirst := stack.first;
	var first Node
	first.item = item;
	first.next = oldFirst;
	stack.first = &first
	stack.length++
}

func (stack *LinkedListStack) pop() string {
	item := (*stack.first).item;
	stack.first = stack.first.next
	stack.length--
	return item;
}

func main() {
	stack := LinkedListStack{0,nil}
	stack.push("a")
	stack.push("b")
	stack.push("c")
	fmt.Println(stack.size())
	temp := stack.pop()
	fmt.Println(temp)
	fmt.Println(stack.size())
}
