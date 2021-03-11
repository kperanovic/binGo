package memory

import (
	"errors"

	"github.com/kperanovic/tombola/internal/common"
)

// GetItem returns a single item entry from memory.
func (m *MemoryStore) GetItem(id int) (*common.Item, error) {
	m.Lock()
	defer m.Unlock()

	if m.itemExists(id) {
		return nil, errors.New("item doesn't exist")
	}

	return m.items[id], nil
}

// GetAllItems will fetch all the saved items from the memory
func (m *MemoryStore) GetAllItems() ([]*common.Item, error) {
	m.Lock()
	defer m.Unlock()

	if len(m.items) == 0 {
		return nil, errors.New("no items saved")
	}

	items := make([]*common.Item, 0)

	for _, item := range m.items {
		items = append(items, item)
	}

	return items, nil
}

// SetItem will save an item in memory storage
func (m *MemoryStore) SetItem(req *common.Item) error {
	m.Lock()
	defer m.Unlock()

	// First check if the player is already saved
	if m.itemExists(req.ID) {
		return errors.New("item already exists")
	}

	// Saving the player
	m.items[req.ID] = req

	// NOTE: the check is not necessary but it is ok to have it in
	// debug mode.
	// Check if the player is saved correctly
	if m.itemExists(req.ID) {
		// debug log that the player is successfuly saved.
		// Also log the structure.
	}

	return nil
}

// DeleteItem will delete an item from memory store
func (m *MemoryStore) DeleteItem(id int) error {
	m.Lock()
	defer m.Unlock()

	// Check if the player exists in the storage
	if !m.itemExists(id) {
		return errors.New("item doesn't exist")
	}

	delete(m.items, id)

	// NOTE: Double checking if the player is actually deleted from
	// the map. This check shouldn't be necessary but, hey its Go.
	if m.itemExists(id) {
		// This is actually a big problem, and the function should
		// return an error. Also, report a bug to the Go community to fix
		// delete func.
	}

	return nil
}

func (m *MemoryStore) itemExists(id int) bool {
	if _, exists := m.items[id]; exists {
		return true
	}

	return false
}
