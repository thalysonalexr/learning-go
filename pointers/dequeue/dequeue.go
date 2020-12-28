// reference: https://github.com/thalysonalexr/data-structures/blob/master/deque/model3.py
package dequeue

import (
	"errors"
	"fmt"
)

const TEMPLATE_FORMAT = "[%d] - %s (%d)\n"

type User struct {
	Id   int
	Name string
	Age  int
}

type Node struct {
	value User
	next  *Node
	prev  *Node
}

type Dequeue struct {
	head   *Node
	tail   *Node
	Length int
}

func NewDequeue() *Dequeue {
	return &Dequeue{}
}

func (q *Dequeue) AddHead(value User) {
	node := Node{
		value: value,
	}
	if q.isEmpty() {
		q.head = &node
		q.tail = &node
	} else {
		node.next = q.head
		q.head.prev = &node
		q.head = &node
	}
	q.Length += 1
}

func (q *Dequeue) AddTail(value User) {
	node := Node{
		value: value,
	}
	if q.isEmpty() {
		q.head = &node
		q.tail = &node
	} else {
		node.prev = q.tail
		q.tail.next = &node
		q.tail = &node
	}
	q.Length += 1
}

func (q *Dequeue) RemoveHead() {
	if q.isEmpty() {
		return
	}
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.head.prev = nil
	}
	q.Length -= 1
}

func (q *Dequeue) RemoveTail() {
	if q.isEmpty() {
		return
	}
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.prev
		q.tail.next = nil
	}
	q.Length -= 1
}

func (q *Dequeue) isEmpty() bool {
	return q.head == nil
}

func (q *Dequeue) search(index int) (User, error) {
	node := q.head
	i := 0
	for node != nil {
		if index == i {
			return node.value, nil
		}
		i += 1
		node = node.next
	}
	return User{}, errors.New("User not found")
}

func (q *Dequeue) SearchBinary(id int) (User, error) {
	if q.isEmpty() {
		return User{}, errors.New("Dequeue is empty")
	}
	first := 0
	last := q.Length - 1
	for first <= last {
		mid := (first + last) / 2
		user, err := q.search(mid)
		if err == nil {
			if id > user.Id {
				first = mid + 1
			} else if id < user.Id {
				last = mid - 1
			} else {
				return user, nil
			}
		}
	}
	return User{}, errors.New("User not found")
}

func (q *Dequeue) PrintHead() {
	if q.isEmpty() {
		return
	}
	node := q.tail
	for node != nil {
		fmt.Printf(TEMPLATE_FORMAT, node.value.Id, node.value.Name, node.value.Age)
		node = node.prev
	}
}

func (q *Dequeue) PrintTail() {
	if q.isEmpty() {
		return
	}
	node := q.head
	for node != nil {
		fmt.Printf(TEMPLATE_FORMAT, node.value.Id, node.value.Name, node.value.Age)
		node = node.next
	}
}
