package framework

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type Deque[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) PushFront(value T) {
	node := &Node[T]{Value: value}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.Next = d.head
		d.head.Prev = node
		d.head = node
	}
	d.size++
}

func (d *Deque[T]) PushBack(value T) {
	node := &Node[T]{Value: value}
	if d.tail == nil {
		d.head = node
		d.tail = node
	} else {
		node.Prev = d.tail
		d.tail.Next = node
		d.tail = node
	}
	d.size++
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.head == nil {
		var zero T
		return zero, false
	}

	value := d.head.Value
	d.head = d.head.Next
	if d.head == nil {
		d.tail = nil
	} else {
		d.head.Prev = nil
	}
	d.size--
	return value, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.tail == nil {
		var zero T
		return zero, false
	}

	value := d.tail.Value
	d.tail = d.tail.Prev
	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.Next = nil
	}
	d.size--
	return value, true
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}
