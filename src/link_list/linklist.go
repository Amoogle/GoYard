package linklist

import (
	"fmt"
)

type LinkNode struct {
	data interface{}
	next *LinkNode
}

func CreateHead(head *LinkNode, array []int) *LinkNode {
	for _, data := range array {
		node := LinkNode{data: data}
		if head != nil {
			node.next = head.next
			head.next = &node
		}
	}
	return head
}

func CreateTail(head *LinkNode, array []int) *LinkNode {
	temp := head
	for _, data := range array {
		node := LinkNode{data: data}
		if head != nil {
			head.next = &node
			head = head.next
		}
	}
	return temp
}

func Travel(head *LinkNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func InertNode(head *LinkNode, index int, data interface{}) bool {
	if index < 0 {
		return false
	} else {
		node := LinkNode{data: data}
		i := 0
		for ; head.next != nil && i < index; head = head.next {
			i++
		}
		if i >= index && head != nil {
			if i == 0 {
				node.next = head
				head.data, node.data = node.data, head.data
			}
			if i > 0 {
				node.next = head.next
				head.next = &node
			}
		} else {
			return false
		}
		return true
	}
}
