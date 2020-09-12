package linklist

type Node struct {
	data interface{}
	next *Node
}

type OneCycleList struct {
	size int
	head *Node
}

// get the head of list

func (o *OneCycleList) GetHead() *Node {
	return o.head
}

func (o *OneCycleList) GetSize() int {
	return o.size
}

func (o *Node) GetData() interface{} {
	return o.data
}

func (o *Node) GetNext() *Node {
	return o.next
}

func (o *OneCycleList) Append(data interface{}) bool {
	node := Node{}
	node.data = data
	if o.GetSize() == 0 {
		o.head = &node
	} else {
		item := o.GetHead()
		for ; item.next != o.GetHead(); item = item.next {
		}
		item.next = &node
	}
	node.next = o.head
	o.size++
	return true

}

func (o *OneCycleList) InsertNext(element *Node, data interface{}) bool {
	if element == nil {
		return false
	}

	node := Node{}
	node.data = data
	node.next = &node
	o.size++
	return true
}

func (o *OneCycleList) Remove(element *Node) interface{} {
	if element == nil {
		return false
	}
	item := o.GetHead()
	for ; item.next != element; item = item.next {
	}
	item.next = element.next
	o.size--
	return element.GetData()
}

func Reverse(head *Node, o *OneCycleList) *OneCycleList {
	if head.next == o.head {
		o.head = head
		return o
	}
	p := Reverse(head.next, o)
	head.next.next = head
	head.next = p.head
	return p
}
