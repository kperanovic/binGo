package memory

import (
	"errors"

	"github.com/kperanovic/tombola/internal/common"
)

// GetPlayer returns a single player entry from memory
func (m *MemoryStore) GetPlayer(id int) (*common.Player, error) {
	m.Lock()
	defer m.Unlock()

	if m.playerExists(id) {
		return nil, errors.New("player doesn't exist")
	}

	return m.players[id], nil
}

// GetAllPlayers will fetch all the saved players from the memory
func (m *MemoryStore) GetAllPlayers() ([]*common.Player, error) {
	m.Lock()
	defer m.Unlock()

	if len(m.players) == 0 {
		return nil, errors.New("no players saved")
	}

	players := make([]*common.Player, 0)

	for _, player := range m.players {
		players = append(players, player)
	}

	return players, nil
}

// SetPlayer will save a player in memory storage
func (m *MemoryStore) SetPlayer(req *common.Player) error {
	m.Lock()
	defer m.Unlock()

	// First check if the player is already saved
	if m.playerExists(req.ID) {
		return errors.New("player already exists")
	}

	// Saving the player
	m.players[req.ID] = req

	// NOTE: the check is not necessary but it is ok to have it in
	// debug mode.
	// Check if the player is saved correctly
	if m.playerExists(req.ID) {
		// debug log that the player is successfuly saved.
		// Also log the structure.
	}

	return nil
}

// DeletePlayer will delete a player from memory store
func (m *MemoryStore) DeletePlayer(id int) error {
	m.Lock()
	defer m.Unlock()

	// Check if the player exists in the storage
	if !m.playerExists(id) {
		return errors.New("player doesn't exist")
	}

	delete(m.players, id)

	// NOTE: Double checking if the player is actually deleted from
	// the map. This check shouldn't be necessary but, hey its Go.
	if m.playerExists(id) {
		// This is actually a big problem, and the function should
		// return an error. Also, report a bug to the Go community to fix
		// delete func.
	}

	return nil
}

func (m *MemoryStore) playerExists(id int) bool {
	if _, exists := m.players[id]; exists {
		return true
	}

	return false
}
