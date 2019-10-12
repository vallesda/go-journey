package main

import (
	"fmt"
)

type Node struct {
	value	int
	next	*Node
}

type NodeList struct {
	start	*Node
	end		*Node
	length	int
}

func (n *NodeList) Add(key int) {
	newNode := &Node{ key, nil }
	if n.start == nil {
		n.start = newNode
		n.end = newNode
	} else {
		n.end.next = newNode
		n.end = n.end.next
	}
}

func (n *NodeList) Remove(key int) {
	if n.start == nil { return }
	if n.start.value == key {
		temp := n.start
		n.start = n.start.next
		*&temp =  nil
	} else {
		current := n.start
		prev := &Node{}
		for current != nil && current.value !=  key {
			prev = current
			current = current.next
		}

		if (current != nil && current.value == key) {
			prev.next = current.next
			current = nil
		}
	}
}

func (n *NodeList) Print() {
	current := n.start
	for current != nil {
		fmt.Printf("%v ->", current.value)
		current = current.next
	}
}

func main () {
	n := NodeList{}
	n.Add(0)
	n.Add(1)
	n.Remove(0)
	n.Add(2)
	n.Add(3)
	n.Remove(-1)
	n.Remove(3)
	n.Print()
}