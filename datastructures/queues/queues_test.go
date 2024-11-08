package queues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Clear(t *testing.T) {
	queue := New[string]()
	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, queue.IsEmpty(), true)

	queue.Enqueue("hello")
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)

	queue.Clear()
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)
}

func TestQueue_GetValues(t *testing.T) {
	queue := New[string]()

	queue.Enqueue("hello", "abc", "xyz")
	// assert.True(t, reflect.DeepEqual(queue.GetValues(), []string{"hello", "abc", "xyz"}))
	assert.Equal(t, queue.GetValues(), []string{"hello", "abc", "xyz"})
}

func TestQueue_Peek(t *testing.T) {
	queue := New[string]()

	queue.Enqueue("hello")
	value, err := queue.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)

	queue.Enqueue("abc")
	value, err = queue.Peek()
	assert.Equal(t, value, "hello")
	assert.Nil(t, err)
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(111, 222)

	value, err := queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, queue.IsEmpty(), false)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, value, 111)

	value, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, queue.IsEmpty(), true)
	assert.Equal(t, queue.Size(), 0)
	assert.Equal(t, value, 222)

	_, err = queue.Dequeue()
	assert.NotNil(t, err)
}

// BenchmarkEnqueue-10    	 8999667	       130.4 ns/op	     115 B/op	       2 allocs/op
func BenchmarkEnqueue(b *testing.B) {
	q := New[string]()
	for i := 0; i < b.N; i++ {
		q.Enqueue("v" + fmt.Sprintf("%d", i))
	}

	b.ReportAllocs()
}

func seedQueue(q *Queue[string], num int) {
	for n := 0; n < num; n++ {
		q.Enqueue("v" + fmt.Sprintf("%d", n))
	}
}
func BenchmarkDequeue(b *testing.B) {
	q := New[string]()
	seedQueue(q, 800000)
	// fmt.Println("init len: ", q.Size())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// q.Enqueue("v" + fmt.Sprintf("%d", i))
		if _, err := q.Dequeue(); err != nil {
			break
		}
	}
	b.StopTimer()

	// b.ReportAllocs()
	// b.ReportMetric()
}
