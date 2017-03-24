package main

import (
	"sync"
	"fmt"
)

func main() {
	q1 := Queue{stack_newest:NewStack(), stack_oldest:NewStack()}
	var q, _type, x int

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&_type)
		if (_type == 1) {
			fmt.Scan(&x)
			q1.push(x)
		} else if (_type == 2) {
			q1.pop()
		} else {
			fmt.Printf("%d\n", q1.front())
		}
	}

}

type Queue struct {
	stack_newest *Stack
	stack_oldest *Stack
}

func (q *Queue) push(data int) {
	q.stack_newest.pushRune(data)
}

func (q *Queue) pop() {
	if (q.stack_oldest.count == 0) {
		for q.stack_newest.count > 0 {
			q.stack_oldest.pushRune(q.stack_newest.popRune())
		}
	}
	q.stack_oldest.popRune()
}

func (q *Queue) front() int{
	if (q.stack_oldest.count == 0) {
		for q.stack_newest.count > 0 {
			q.stack_oldest.pushRune(q.stack_newest.popRune())
		}
	}
	return q.stack_oldest.peekRune()
}

type stackNode struct {
	data int
	next *stackNode
}

type Stack struct {
	top   *stackNode
	count int
	lock  *sync.Mutex
}

func NewNode(item int) *stackNode {
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

func (s *Stack) peek() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.top
	return n.data
}

func (s *Stack) push(data int) {
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

func (s *Stack) pop() int {
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
