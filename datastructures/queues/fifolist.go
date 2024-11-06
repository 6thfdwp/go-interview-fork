package queues

import "container/list"

type FIFOQueue struct {
	items *list.List
}

func NewFIFOQueue() *FIFOQueue {
	return &FIFOQueue{items: list.New()}
}

func (t *FIFOQueue) Size() int {
	return t.items.Len()
}

func (t *FIFOQueue) Enqueue(vals ...string) {
	// push one to many to the end
	for _, v := range vals {
		t.items.PushBack(v)
	}
}

func (t *FIFOQueue) Dequeue() string {
	if t.Size() == 0 {
		return ""
	}
	// pop from the start
	e := t.items.Front()
	t.items.Remove(e)
	return e.Value.(string)
}
