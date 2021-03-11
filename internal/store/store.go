package store

import "github.com/kperanovic/tombola/internal/common"

type Store interface {
	GetPlayer(id int) (*common.Player, error)
	GetAllPlayers() ([]*common.Player, error)
	SetPlayer(*common.Player) error
	DeletePlayer(id int) error

	GetItem(id int) (*common.Item, error)
	GetAllItems() ([]*common.Item, error)
	SetItem(*common.Item) error
	DeleteItem(id int) error
}
