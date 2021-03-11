package common

import "time"

type Player struct {
	ID           int
	FirstName    string
	LastName     string
	IsEmployee   bool
	isActive     bool
	RegisteredAt time.Time
}

type Item struct {
	ID           int
	Name         string
	Description  string
	IsAvailable  bool
	RegisteredAt time.Time
}
