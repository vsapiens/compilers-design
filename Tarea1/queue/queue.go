package queue

import (
	"../node"
)

//Queue The structure of FIFO Order
type Queue struct {
	size  int
	final *node.Node
}

//NewQueue Initializes new queue
func NewQueue(newNode *node.Node) *Queue {
	q := Queue{size: 0, final: newNode}
	return &q
}

//Front Returns the front value of the queue
func (self *Queue) Front() interface{} {
	return self.final.GetNext().GetData()
}

//Push Adds in a fifo manner any type of data stores
func (self *Queue) Push(data interface{}) {
	if self.final == nil {
		self.final = node.NewNode(data, nil)
		self.final.SetNext(self.final)
	} else {
		self.final.SetNext(node.NewNode(data, self.final.GetNext()))
		self.final = self.final.GetNext()
	}
	self.size++
}

//Size Returns the size of the queue
func (self *Queue) Size() int {
	return self.size
}

//Pop Deletes in a fifo manner any type of data stores
func (self *Queue) Pop() {
	curr := self.final.GetNext()
	if self.final.GetNext() == nil {
		self.final = nil
	} else {
		self.final.SetNext(curr.GetNext())
	}
	self.size--
}

//Empty the queue
func (self *Queue) Empty() {
	self = nil
}
