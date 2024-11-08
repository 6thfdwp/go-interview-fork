package lrux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRURecentPut(t *testing.T) {
	t.Run("should put new listening to the front", func(t *testing.T) {
		// The new one is like MRU, make existing ones become LRU
		lru := NewLRURecentSongs(3)
		lru.Put(SongEntry{Id: 1, Name: "Cannot stop"})
		lru.Put(SongEntry{Id: 2, Name: "Cabron"})
		lru.Put(SongEntry{Id: 3, Name: "Tear"}) // new listen
		assert.Equal(t, 3, lru.Len())

		first := lru.PeekFirst()
		assert.Equal(t, "Tear", first.Name)

	})

	t.Run("should move to the front when existing song is played again", func(t *testing.T) {
		lru := NewLRURecentSongs(3)
		lru.Put(SongEntry{Id: 1, Name: "Cannot stop"})
		lru.Put(SongEntry{Id: 2, Name: "Cabron"})

		first := lru.PeekFirst()
		assert.Equal(t, "Cabron", first.Name)

		lru.Put(SongEntry{Id: 1, Name: "Cannot stop"})
		assert.Equal(t, 2, lru.Len())
		first = lru.PeekFirst()
		assert.Equal(t, "Cannot stop", first.Name)
	})

	t.Run("should evict the last one if reaching cap when putting new", func(t *testing.T) {
		lru := NewLRURecentSongs(2)
		lru.Put(SongEntry{Id: 1, Name: "Cannot stop"})
		lru.Put(SongEntry{Id: 2, Name: "Cabron"})
		assert.Equal(t, 2, lru.Len())
		el := lru.PeekLast()
		assert.Equal(t, "Cannot stop", el.Name)

		lru.Put(SongEntry{Id: 3, Name: "Scar Tissue"})
		assert.Equal(t, 2, lru.Len())

		firstSong := lru.PeekFirst()
		assert.Equal(t, "Scar Tissue", firstSong.Name)
		lastSong := lru.PeekLast()
		assert.Equal(t, "Cabron", lastSong.Name)
	})
}

func TestLRUGet(t *testing.T) {
	t.Run("should return song if in the LRU", func(t *testing.T) {
		lru := NewLRURecentSongs(2)
		song1 := SongEntry{Id: 1, Name: "Cannot stop"}
		song2 := SongEntry{Id: 2, Name: "Cabron"}
		lru.Put(song1)
		lru.Put(song2)

		res, err := lru.Get(1)
		assert.Equal(t, song1, res)
		assert.NoError(t, err)

		first := lru.PeekFirst()
		assert.Equal(t, song2, first)
	})

	t.Run("should return error if not found in the LRU", func(t *testing.T) {
		lru := NewLRURecentSongs(2)
		song1 := SongEntry{Id: 1, Name: "Cannot stop"}
		song2 := SongEntry{Id: 2, Name: "Cabron"}
		lru.Put(song1)
		lru.Put(song2)

		_, err := lru.Get(3)
		assert.Error(t, err)
	})
}
