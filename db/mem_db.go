package db

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrLeaderboardNotFound = errors.New("leaderboard not found")
	ErrLeaderboardInvalidFieldName = errors.New("invalid field name")
)

var db = make(map[uuid.UUID]*Leaderboard)

type MemDB struct{}

func NewMemDB() DB {
	return &MemDB{}
}

func (d *MemDB) CreateLeaderboard(name string, fieldNames []string) *Leaderboard {
	leaderboard := NewLeaderboard(name, fieldNames)

	db[leaderboard.ID] = leaderboard

	return leaderboard
}

func (d *MemDB) ReadLeaderboard(id uuid.UUID) *Leaderboard {
	return db[id]
}

func (d *MemDB) CreateEntry(
	leaderboardId uuid.UUID,
	name string,
	fieldValues map[string]int,
) (*Entry, error) {
	leaderboard := d.ReadLeaderboard(leaderboardId)
	if leaderboard == nil {
		return nil, ErrLeaderboardNotFound
	}

	for fieldName := range fieldValues {
		if !leaderboard.HasFieldName(fieldName) {
			// TODO: Return the invalid field name.
			return nil, ErrLeaderboardInvalidFieldName
		}
	}

	entry := NewEntry(name, fieldValues)
	leaderboard.Entries = append(leaderboard.Entries, entry)

	return entry, nil
}
