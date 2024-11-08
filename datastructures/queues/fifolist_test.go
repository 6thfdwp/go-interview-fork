package queues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIFOList(t *testing.T) {
	t.Run("should push to end of the Q", func(t *testing.T) {
		q := NewFIFOQueue()
		q.Enqueue("v1", "v4", "v5")
		// q.Enqueue("v5")

		assert.Equal(t, 3, q.Size())
	})

	t.Run("should dequeue from the front of the Q", func(t *testing.T) {
		q := NewFIFOQueue()
		q.Enqueue("v1", "v5")
		r := q.Dequeue()

		assert.Equal(t, 1, q.Size())
		assert.Equal(t, "v1", r)
	})
}

// BenchmarkFIFOList-10    	 6288309	       190.4 ns/op	      88 B/op	       4 allocs/op
func BenchmarkFIFOList(b *testing.B) {
	q := NewFIFOQueue()
	for i := 0; i < b.N; i++ {
		q.Enqueue("v" + fmt.Sprintf("%d", i))
	}

	// b.ReportAllocs()
}
