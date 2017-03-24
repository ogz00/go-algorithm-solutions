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
		s := NewRuneStack()
		for _, bracket := range brackets {

			switch (bracket) {
			case '{':
				s.pushRune('}')
				break
			case '(':
				s.pushRune(')')
				break
			case '[':
				s.pushRune(']')
				break
			default :
				if (s.isEmptyRune() || bracket != s.popRune()) {
					return false
				}
			}

		}
		return s.isEmptyRune();
	}

}

type runeNode struct {
	data rune
	next *runeNode
}

type Stack struct {
	top   *runeNode
	count int
	lock  *sync.Mutex
}

func runeNode(item rune) *runeNode {
	return &runeNode{
		data:item,
	}
}

func NewRuneStack() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

func (s *Stack) isEmptyRune() bool {
	return s.top == nil
}

func (s *Stack) peekRune() rune {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.top
	return n.data
}

func (s *Stack) pushRune(data rune) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := runeNode(data)
	if s.top == nil {
		s.top = n
	} else {
		n.next = s.top
		s.top = n
	}
	s.count++
}

func (s *Stack) popRune() rune {
	s.lock.Lock()
	defer s.lock.Unlock()

	var n *runeNode
	if (s.top != nil) {
		n = s.top
		s.top = n.next
		s.count --
	}

	return n.data

}
