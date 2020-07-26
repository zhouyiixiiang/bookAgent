package basic

type Node struct {
	Data string
	Next *Node
}

type QueueLink struct {
	rear  *Node
	front *Node
}

type LinkQueue interface {
	Lengh() int
	Enqueue(value interface{})
	Dequeue() (value interface{})
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) Length() int {
	next := qlk.front
	length := 0
	for next.Next != nil {
		next = next.Next
		length++
	}
	return length
}
func (qlk *QueueLink) Enqueue(value string) {
	newNode := &Node{Data: value, Next: nil}
	if qlk.front == nil {
		qlk.front = newNode
		qlk.rear = newNode
	} else {
		qlk.rear.Next = newNode
		qlk.rear = qlk.rear.Next
	}
}
func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newNode := qlk.front
	if qlk.front == qlk.rear {
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.Next
	}
	return newNode.Data
}
