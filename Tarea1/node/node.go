package node

//Node The structure a container holding a value
type Node struct {
	val  interface{}
	next *Node
}

//SetData Sets the data of any type to the node
func (self *Node) SetData(val interface{}) {
	self.val = val
}

//GetData Gets the data of any type to the node
func (self *Node) GetData() interface{} {
	return self.val
}

//GetNext Gets the next Node
func (self *Node) GetNext() *Node {
	return self.next
}

//SetNext Sets the next Node
func (self *Node) SetNext(newNode *Node) {
	self.next = newNode
}

//NewNode Creates new Node
func NewNode(data interface{}, newNode *Node) *Node {
	t := Node{val: data, next: newNode}
	return &t
}
