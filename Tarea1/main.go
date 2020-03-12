package main

import (
	"fmt"

	"./hash"
	"./node"
	"./queue"
	"./stack"
)

func main() {
	//Stack Works
	var node1 node.Node
	s := stack.NewStack(&node1)
	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(1)
	sizeStack := s.Size()
	for ; sizeStack > 0; sizeStack-- {
		fmt.Printf("%v \t", s.Top())
		s.Pop()
	}
	fmt.Printf("\n")

	//Queue works
	var q1 queue.Queue
	q1.Push(1)
	q1.Push(2)
	q1.Push(3)
	q1.Push(4)
	q1.Push(5)
	q1.Push(6)
	q1.Push(7)

	sizeQueue := q1.Size()
	for ; sizeQueue > 0; sizeQueue-- {
		fmt.Printf("%v \t", q1.Front())
		q1.Pop()
	}
	fmt.Printf("\n")

	//Hash works
	var h hash.Map
	h.Insert(1, "Soy el 1")
	h.Insert(21, 2)
	h.Insert(222, 3)
	h.Insert(12, 4)
	fmt.Printf("%v \t", h.Get(12))
	fmt.Printf("%v \t", h.Get(1))
	fmt.Printf("\n")
}
