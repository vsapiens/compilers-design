package hash

const HASH_NUMBER = 7

type HashNode struct {
	index int
	data  interface{}
	next  *HashNode
}

type Map struct {
	hash [HASH_NUMBER]*HashNode
}

func hash(index int) int {
	return index % HASH_NUMBER
}

//Insert Inserts a hashed node in the map
func (self *Map) Insert(index int, data interface{}) {

	hashed_index := hash(index)

	if self.hash[hashed_index] == nil {
		self.hash[hashed_index] = &HashNode{index: index, data: data, next: nil}
	} else {
		curr := self.hash[hashed_index]
		//Find the next empty space
		for curr.next != nil {
			if curr.index == index {
				return
			}
			curr = curr.next
		}
		//If the index already exists
		if curr.index == index {
			return
		}
		curr.next = &HashNode{index: index, data: data, next: nil}
	}
}

//Get Gets a hashed node in the map
func (self *Map) Get(index int) interface{} {
	hashed_index := hash(index)

	if self.hash[hashed_index] != nil {
		curr := self.hash[hashed_index]

		for curr != nil {
			if curr.index == index {
				return curr.data
			}
			curr = curr.next
		}
	}

	return nil
}

//Remove Remove a hashed node in the map
func (self *Map) Remove(index int) {
	hashed_index := hash(index)

	if self.hash[hashed_index] != nil {

		curr := self.hash[hashed_index]

		//If the index already exists
		if curr.index == index {
			self.hash[hashed_index] = curr.next
			return
		}
		//Find the next empty space
		for curr.next != nil {
			if curr.next.index == index {
				curr.next = curr.next.next
				return
			}
			curr = curr.next
		}

	}

}
