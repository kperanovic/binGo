package common

import "time"

// Selection is a struct that saves
// the random selection made by the engine.
type Selection struct {
	Player   *Player
	Item     *Item
	ChosenAt time.Time
}

// Combinations is a structure that saves the list
// of all the players and items that were selected.
// Players and items are bound to each other.
type Combinations struct {
	Selections []*Selection
}
