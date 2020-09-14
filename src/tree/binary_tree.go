package tree

import "goyard/src/linklist"

type Node struct {
	data  interface{}
	right *linklist.LinkNode
	left  *linklist.LinkNode
}
