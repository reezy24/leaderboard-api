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

package db

import "github.com/google/uuid"

type DB interface {
	// Leaderboard operations.
	CreateLeaderboard(name string, columns []string) *Leaderboard
	ReadLeaderboard(id uuid.UUID) *Leaderboard
	// Entry operations.
	CreateEntry(leaderboardId uuid.UUID, name string, fieldValues map[string]int) (*Entry, error)
}

type Leaderboard struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	FieldNames []string  `json:"fieldNames"`
	Entries    []*Entry  `json:"entries,omitempty"`
}

type Entry struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	FieldValues map[string]int `json:"fieldValues"`
}

func NewLeaderboard(name string, fieldNames []string) *Leaderboard {
	return &Leaderboard{
		ID:         uuid.New(),
		Name:       name,
		FieldNames: fieldNames,
	}
}

func NewEntry(name string, fieldValues map[string]int) *Entry {
	return &Entry{
		ID:          uuid.New(),
		Name:        name,
		FieldValues: fieldValues,
	}
}

func (l *Leaderboard) HasFieldName(name string) bool {
	for _, fname := range l.FieldNames {
		if fname == name {
			return true
		}
	}

	return false
}
