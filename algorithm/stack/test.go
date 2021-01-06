package main

import "fmt"

func main() {
	stack := ResizingArrayStack{[]string{"a"}, 0}
	stack.push("a")
	stack.push("b")
	stack.push("c")
	fmt.Println(stack.size());
	temp := stack.pop();
	fmt.Println(temp);
	fmt.Println(stack.size());
}
