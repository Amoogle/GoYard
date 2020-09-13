package main

import (
	"fmt"
	linklist "goyard/src/link_list"
)

func main() {
	//node := linklist.Node{}
	cycleLink := linklist.OneCycleList{}
	cycleLink.Append(2)
	array := []int{1, 2, 3, 4, 5}
	//linklist.Reverse(&node, &cycleLink)

	link1 := linklist.LinkNode{}
	head1 := linklist.CreateHead(&link1, array)
	linklist.InsertNode(head1, 8, 99)
	linklist.Travel(&link1)

	link2 := linklist.LinkNode{}
	head2 := linklist.CreateTail(&link2, array)
	linklist.Travel(head2)
	linklist.InsertNode(head2, 5, 99)
	linklist.Travel(&link2)

	rlink := linklist.ReverseLink(&link2)
	linklist.Travel(rlink)

	o := new(linklist.LinkNode)
	linklist.RReverse(o, rlink)
	linklist.Travel(o)
	fmt.Println("hello go word")
}
