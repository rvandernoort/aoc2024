package framework

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type Deque struct {
	head *Node
	tail *Node
	size int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) PushFront(value interface{}) {
	node := &Node{Value: value}
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

func (d *Deque) PushBack(value interface{}) {
	node := &Node{Value: value}
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

func (d *Deque) PopFront() interface{} {
	if d.head == nil {
		return nil
	}

	value := d.head.Value
	d.head = d.head.Next
	if d.head == nil {
		d.tail = nil
	} else {
		d.head.Prev = nil
	}
	d.size--
	return value
}

func (d *Deque) PopBack() interface{} {
	if d.tail == nil {
		return nil
	}

	value := d.tail.Value
	d.tail = d.tail.Prev
	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.Next = nil
	}
	d.size--
	return value
}

func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}
