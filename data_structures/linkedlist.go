package linkedlist

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Element generic.Type

type Node struct {
	val  Element
	next *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}
