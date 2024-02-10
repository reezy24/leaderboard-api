package db

import "github.com/google/uuid"

var db = make(map[uuid.UUID]*Leaderboard)

type DB struct{}

type Leaderboard struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Columns []string  `json:"columns"`
	Entries []*Entry  `json:"entries,omitempty"`
}

type Entry struct{}

/*
example leaderboard:

player | points | kills | money
-------|--------|-------|------
reez   | 69     | 0  	| 0
hayley | 420 	| 69	| 69420

problem:
a leaderboard can have n number of columns

how do we represent a leaderboard and a leaderboard entry
such that columns and it's values are parameters
*/

func NewLeaderboardDB() *DB {
	return &DB{}
}

func NewLeaderboard(name string, columns []string) *Leaderboard {
	return &Leaderboard{
		ID:      uuid.New(),
		Name:    name,
		Columns: columns,
	}
}

func (l *DB) CreateLeaderboard(name string, columns []string) *Leaderboard {
	leaderboard := NewLeaderboard(name, columns)

	db[leaderboard.ID] = leaderboard

	return leaderboard
}

func (l *DB) ReadLeaderboard(id uuid.UUID) *Leaderboard {
	return db[id]
}
