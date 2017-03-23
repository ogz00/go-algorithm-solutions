package main

import (
	"fmt"
	"sync"
)

func main() {

	var turn int
	fmt.Scan(&turn)

	for i := 0; i < turn; i++ {
		var brackets string
		fmt.Scan(&brackets)
		if (isBalanced(brackets)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func isBalanced(expression string) bool {

	// Must be even
	if ((len(expression) & 1) == 1) {
		return false
	} else {
		brackets := [] rune(expression)
		s := NewStack()
		for _, bracket := range brackets {

			switch (bracket) {
			case '{':
				s.push('}')
				break
			case '(':
				s.push(')')
				break
			case '[':
				s.push(']')
				break
			default :
				if (s.isEmpty() || bracket != s.pop()) {
					return false
				}
			}

		}
		return s.isEmpty();
	}

}

type stackNode struct {
	data rune
	next *stackNode
}

type Stack struct {
	top   *stackNode
	count int
	lock  *sync.Mutex
}

func NewNode(item rune) *stackNode {
	return &stackNode{
		data:item,
	}
}

func NewStack() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) peek() rune {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.top
	return n.data
}

func (s *Stack) push(data rune) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := NewNode(data)
	if s.top == nil {
		s.top = n
	} else {
		n.next = s.top
		s.top = n
	}
	s.count++
}

func (s *Stack) pop() rune {
	s.lock.Lock()
	defer s.lock.Unlock()

	var n *stackNode
	if (s.top != nil) {
		n = s.top
		s.top = n.next
		s.count --
	}

	return n.data

}
