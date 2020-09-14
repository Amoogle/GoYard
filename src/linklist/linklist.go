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
	head.data = array[0]
	for _, data := range array[1:] {
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

// 插入指定节点
func InsertNode(head *LinkNode, index int, data interface{}) bool {
	// 函数体内head指针为实参的一个拷贝
	if index < 0 {
		return false
	}
	node := LinkNode{data: data}

	// 插入节点到链表首位
	if index == 0 {
		node.next = head.next
		head.next = &node
		head.data, node.data = node.data, head.data
		return true
	}

	i := 0
	for ; head.next != nil && i < index-1; head = head.next {
		i++
	}

	// 插入点超出链表最大可插入点
	if head.next == nil && i < index-1 {
		return false
	}

	// 插入到中间节点
	if head.next != nil {
		node.next = head.next
		head.next = &node
		return true
	}

	// 插入到尾节点
	if head.next == nil {
		head.next = &node
		return true
	}
	return false
}

func ReverseLink(head *LinkNode) *LinkNode {
	cur := head.next
	next := head.next.next
	head.next = nil
	for ; next != nil; next = next.next {
		cur.next = head
		head = cur
		cur = next
	}
	cur.next = head
	return cur
}

// 递归逆转 传入o 作为逆转后的链头
func RReverse(o *LinkNode, head *LinkNode) *LinkNode {
	if head.next == nil {
		o.next = head
		return head
	}
	node := RReverse(o, head.next)
	head.next = nil
	node.next = head
	return head
}
