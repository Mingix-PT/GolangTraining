package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](val T) *List[T] {
	l := new(List[T])
	l.val = val
	l.next = nil
	return l
}

func (l *List[T]) PushFront(val T) {
	l.next = &List[T] {
		val: l.val,
		next: l.next,
	}
	l.val = val
}

func (l *List[T]) PopFront() T {
	val := l.val
	l.val = l.next.val
	l.next = l.next.next
	return val
}

func (l *List[T]) PushBack(val T) {
	for l.next != nil {
		l = l.next
	}
	l.next = &List[T] {
		val: val,
		next: nil,
	}
}

func (l *List[T]) PopBack() T {
	for l.next.next != nil {
		l = l.next
	}
	val := l.next.val
	l.next = nil
	return val
}

func (l *List[T]) Len() int {
	var len int
	for l != nil {
		len++
		l = l.next
	}
	return len
}

func (l *List[T]) At(i int) T {
	for i > 0 {
		l = l.next
		i--
	}
	return l.val
}

func (l *List[T]) Reverse() *List[T] {
	var prev *List[T]
	for l != nil {
		next := l.next
		l.next = prev
		prev = l
		l = next
	}
	return prev
}

func (l *List[T]) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%v -> ", l.val)
		l = l.next
	}
	return s
}

func main() {
	l := NewList(0)
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	fmt.Println(l)

	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	fmt.Println(l)

	fmt.Println(l.PopFront())
	fmt.Println(l)

	fmt.Println(l.PopBack())  
	fmt.Println(l)            

	fmt.Println(l.Len())

	fmt.Println(l.At(1))

	l = l.Reverse()
	fmt.Println(l)
}