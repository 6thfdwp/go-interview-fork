package queues

import (
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
