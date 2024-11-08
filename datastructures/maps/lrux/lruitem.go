package lrux

import (
	"container/list"
	"errors"
)

type SongEntry struct {
	Id   int
	Name string
}

type LRURecentSongs struct {
	capacity    int
	doubleLList *list.List
	itemsById   map[int]*list.Element
}

func NewLRURecentSongs(cap int) *LRURecentSongs {
	return &LRURecentSongs{
		capacity:    cap,
		doubleLList: list.New(),
		// itemsById:   make(map[int]SongEntry),
		// map holds the node, SongEntry is the value in the node
		itemsById: make(map[int]*list.Element),
	}
}

func (t *LRURecentSongs) Len() int {
	return t.doubleLList.Len()
	// return len(t.itemsById)
}

// Put is called when a new song is played
func (t *LRURecentSongs) Put(song SongEntry) {
	if node, ok := t.itemsById[song.Id]; ok {
		// move existing to the front
		t.doubleLList.MoveToFront(node)
		return
	}

	if t.Len() >= t.capacity {
		// evict the last (LRU) from the end
		// also delete the map key
		lastEl := t.doubleLList.Back()
		lastKey := lastEl.Value.(SongEntry).Id
		t.doubleLList.Remove(lastEl)
		delete(t.itemsById, lastKey)
	}
	newNode := t.doubleLList.PushFront(song)
	t.itemsById[song.Id] = newNode
}

func (t *LRURecentSongs) Get(key int) (SongEntry, error) {
	entry, ok := t.itemsById[key]
	if !ok {
		return SongEntry{}, errors.New("no such song")
	}

	return entry.Value.(SongEntry), nil

}

func (t *LRURecentSongs) PeekFirst() SongEntry {
	return t.doubleLList.Front().Value.(SongEntry)
}
func (t *LRURecentSongs) PeekLast() SongEntry {
	return t.doubleLList.Back().Value.(SongEntry)
}
