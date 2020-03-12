package stack

import (
	"../node"
)

//Stack The structure of LIFO Order
type Stack struct {
	size int
	pile *node.Node
}

//NewStack Initializes new Stack
func NewStack(newNode *node.Node) *Stack {
	s := Stack{size: 0, pile: newNode}
	return &s
}

//Push Adds in a lifo manner any type of data stores
func (self *Stack) Push(data interface{}) {
	self.pile = node.NewNode(data, self.pile)
	self.size++
}

//Pop Deletes in a lifo manner any type of data stores
func (self *Stack) Pop() {
	self.pile = self.pile.GetNext()
	self.size--
}

//Top Returns the top value of the stack
func (self *Stack) Top() interface{} {
	return self.pile.GetData()
}

//Size Returns the size of the stack
func (self *Stack) Size() int {
	return self.size
}

//Empty the stack
func (self *Stack) Empty() {
	self = nil
}
