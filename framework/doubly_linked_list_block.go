package framework

type BlockNode struct {
	Value FsBlock
	Next  *BlockNode
	Prev  *BlockNode
}

type BlockDeque struct {
	head *BlockNode
	tail *BlockNode
	size int
}

func NewBlockDeque() *BlockDeque {
	return &BlockDeque{}
}

func (d *BlockDeque) PushFront(value FsBlock) {
	BlockNode := &BlockNode{Value: value}
	if d.head == nil {
		d.head = BlockNode
		d.tail = BlockNode
	} else {
		BlockNode.Next = d.head
		d.head.Prev = BlockNode
		d.head = BlockNode
	}
	d.size++
}

func (d *BlockDeque) PushBack(value FsBlock) {
	BlockNode := &BlockNode{Value: value}
	if d.tail == nil {
		d.head = BlockNode
		d.tail = BlockNode
	} else {
		BlockNode.Prev = d.tail
		d.tail.Next = BlockNode
		d.tail = BlockNode
	}
	d.size++
}

func (d *BlockDeque) PopFront() (FsBlock, bool) {
	if d.head == nil {
		var zero FsBlock
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

func (d *BlockDeque) PopBack() (FsBlock, bool) {
	if d.tail == nil {
		var zero FsBlock
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

type FsBlock struct {
	Size       int
	StartIndex int
	Number     int
}

func (d *BlockDeque) SearchFirstGap(size int) (*FsBlock, int, bool) {
	current := d.head
	index := 1
	for current != nil {
		if current.Value.Number == -1 {
			if current.Value.Size >= size {
				return &current.Value, index, true
			}
		}
		current = current.Next
		index++
	}
	return &FsBlock{}, index, false
}

func (d *BlockDeque) InsertAt(index int, value FsBlock) {
	if index < 0 || index > d.size {
		return
	}

	newNode := &BlockNode{Value: value}

	if index == 0 {
		d.PushFront(value)
		return
	}

	if index == d.size {
		d.PushBack(value)
		return
	}

	current := d.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	newNode.Prev = current.Prev
	newNode.Next = current
	current.Prev.Next = newNode
	current.Prev = newNode
	d.size++
}

func (d *BlockDeque) FirstValue() int {
	if d.head == nil {
		return 999999999
	}
	current := d.head
	for current != nil {
		if current.Value.Number != -1 {
			return current.Value.Number
		}
		current = current.Next
	}
	return 9999999
}

func (d *BlockDeque) LastValue() int {
	return d.tail.Value.Number
}

func (d *BlockDeque) Size() int {
	return d.size
}

func (d *BlockDeque) IsEmpty() bool {
	return d.size == 0
}
