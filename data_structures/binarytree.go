package binarytree

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Element generic.Type

type Node struct {
	val   Element
	left  *Node
	right *Node
}

type BinaryTree struct {
	root Node
	lock sync.RWMutex
}
