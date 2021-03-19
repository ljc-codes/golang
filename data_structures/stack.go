package stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Element generic.Type

type Stack struct {
	elements []Element
	lock     sync.RWMutex
}

func (s *Stack) push(e Element) {
	s.lock.Lock()
	s.elements = append(s.elements, e)
	s.lock.Unlock()
}

func (s *Stack) pop() *Element {
	s.lock.Lock()
	size := len(s.elements)
	elem := s.elements[size-1]
	s.elements = s.elements[:size-1]
	s.lock.Unlock()
	return &elem
}
