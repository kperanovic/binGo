package memory

import (
	"sync"

	"github.com/kperanovic/tombola/internal/common"
)

// MemoryStore is a skeleton structure for memory storage
type MemoryStore struct {
	sync.Mutex
	players map[int]*common.Player
	items   map[int]*common.Item
}

//NewMemoryStore creates a new MemoryStore instance
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		players: make(map[int]*common.Player),
		items:   make(map[int]*common.Item),
	}
}
