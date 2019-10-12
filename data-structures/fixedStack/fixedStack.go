package main

import "fmt"

type FixedStack struct {
	size int
	top int
	items []int 
}

func (fs *FixedStack) New(size int) *FixedStack {
	fs.size = size
	fs.top = -1
	fs.items = make([]int, size)
	return fs
}

func (fs *FixedStack) Push(v int) {
	if !fs.IsFull() {
		fs.top++
		fs.items[fs.top] = v
	} else {
		fmt.Println("Stack is Full")
	}
}

func (fs *FixedStack) Pop() int {
	v := -1
	if !fs.IsEmpty() {
		v = fs.items[fs.top]
		fs.top--
	}
	return v
}

func (fs *FixedStack) Top() int {
	v := -1
	if !fs.IsEmpty() {
		v = fs.items[fs.top]
	}
	return v
}

func (fs *FixedStack) Size() int {
	return fs.top + 1
}

func (fs *FixedStack) IsEmpty() bool {
	return fs.top == -1
}

func (fs *FixedStack) IsFull() bool {
	return fs.top == fs.size - 1
}

func main () {
	var stack FixedStack
	stack.New(2)
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Size())
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	stack.Top()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
}
