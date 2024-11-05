package hashsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSetStr(t *testing.T) {
	t.Run("should add without duplicates", func(t *testing.T) {
		hset := NewHashSetStr()
		hset.Add("a")
		hset.Add("b")

		assert.Equal(t, 2, hset.Size())

		hset.Add("a")
		assert.Equal(t, 2, hset.Size())
		t.Logf("current set %v", hset)
	})

	t.Run("should check if contains", func(t *testing.T) {
		hset := NewHashSetStr()
		hset.Add("a")
		hset.Add("b")

		assert.True(t, hset.Contains("a"))
		assert.True(t, hset.Contains("b"))
		assert.False(t, hset.Contains("c"))
	})
}
