package engine

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/kperanovic/tombola/internal/common"
	"github.com/kperanovic/tombola/internal/store"
	"go.uber.org/zap"
)

// Engine is in charge of the whole application logic.
type Engine struct {
	log *zap.Logger
	sync.Mutex
	ctx     context.Context
	store   store.Store
	storage *common.Combinations
}

// NewEngine returns a new Engine instance.
func NewEngine(ctx context.Context, log *zap.Logger, store store.Store) *Engine {
	return &Engine{
		ctx:   ctx,
		store: store,
		storage: &common.Combinations{
			Selections: make([]*common.Selection, 0),
		},
	}
}

// Randomize will randomly choose 1 player and 1 item and return it in the
// *common.Selection struct
func (e *Engine) Randomize() (*common.Selection, error) {
	// Randomly choose a player
	players, err := e.store.GetAllPlayers()
	if err != nil {
		e.log.Error("error fetching players", zap.Error(err))
	}
	player := players[rand.Intn(len(players))]

	// Randomly choose an item
	items, err := e.store.GetAllItems()
	if err != nil {
		e.log.Error("error fetching items", zap.Error(err))
	}
	item := items[rand.Intn(len(players))]

	res := &common.Selection{
		Player:   player,
		Item:     item,
		ChosenAt: time.Now(),
	}

	return res, nil
}

// VerifySelection will validate if the randomized selection is
// correct, save the selection in history, and delete both the player and
// the item from the store.
func (e *Engine) VerifySelection(selection *common.Selection) error {
	// First check if the player and the item were eligible for selection
	if _, err := e.store.GetPlayer(selection.Player.ID); err != nil {
		e.log.Error("player was not eligible to obtain an item", zap.Error(err))

		return err
	}

	if _, err := e.store.GetItem(selection.Item.ID); err != nil {
		e.log.Error("item is not obtainable", zap.Error(err))

		return err
	}

	// Delete both the player and the item from the list.
	if err := e.store.DeletePlayer(selection.Player.ID); err != nil {
		e.log.Error("error deleting player from the list", zap.Error(err))
	}

	if err := e.store.DeleteItem(selection.Item.ID); err != nil {
		e.log.Error("error deleting item from the list", zap.Error(err))
	}

	e.log.Info("verificaton successful", zap.Int("playerID", selection.Player.ID), zap.Int("itemID", selection.Item.ID))

	// Save the selection
	e.saveSelection(selection)

	return nil
}

// FetchAllPlayers will return the list of all the players.
func (e *Engine) FetchAllPlayers() ([]*common.Player, error) {
	return e.store.GetAllPlayers()
}

// FetchAllItems will return the list of all the items.
func (e *Engine) FetchAllItems() ([]*common.Item, error) {
	return e.store.GetAllItems()
}

func (e *Engine) saveSelection(selection *common.Selection) {
	e.storage.Selections = append(e.storage.Selections, selection)
}
