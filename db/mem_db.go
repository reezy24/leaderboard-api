package db

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

var (
	ErrLeaderboardNotFound = errors.New("leaderboard not found")
	ErrLeaderboardInvalidFieldName = errors.New("invalid field name")
	ErrLeaderboardHasNoEntries = errors.New("leaderboard has no entries")
	ErrEntryNotFound = errors.New("entry not found")
	ErrEntryInvalidSortKey = errors.New("invalid sort key")
)

var db = make(map[uuid.UUID]*Leaderboard)

type MemDB struct{}

func NewMemDB() DB {
	return &MemDB{}
}

func (d *MemDB) CreateLeaderboard(
	name string,
	fieldNames []string,
	entries map[uuid.UUID]*Entry,
) *Leaderboard {
	e := entries
	if e == nil {
		e = make(map[uuid.UUID]*Entry)
	}

	leaderboard := NewLeaderboard(name, fieldNames, e)

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

	fv := fieldValues
	if fv == nil {
		fv = make(map[string]int)
	}

	entry := NewEntry(name, fv)
	leaderboard.Entries[entry.ID] = entry

	return entry, nil
}

func (d *MemDB) UpdateEntry(
	leaderboardId uuid.UUID,
	entryId uuid.UUID,
	fieldValues map[string]int,
) (*Entry, error) {
	leaderboard := d.ReadLeaderboard(leaderboardId)
	if leaderboard == nil {
		return nil, ErrLeaderboardNotFound
	}

	entry := leaderboard.Entries[entryId]
	if entry == nil {
		return nil, ErrEntryNotFound
	}

	for fieldName := range fieldValues {
		if !leaderboard.HasFieldName(fieldName) {
			// TODO: Return the invalid field name.
			return nil, ErrLeaderboardInvalidFieldName
		}
	}

	for fieldName, v := range fieldValues {
		entry.FieldNamesToValues[fieldName] = v	
	}

	return entry, nil
}

func (d *MemDB) ReadEntries(leaderboardId uuid.UUID, sortKey, order string) ([]*Entry, error) {
	leaderboard := d.ReadLeaderboard(leaderboardId)
	if leaderboard == nil {
		return nil, ErrLeaderboardNotFound
	}

	if len(leaderboard.Entries) == 0 {
		return nil, ErrLeaderboardHasNoEntries
	}

	if sortKey != "" && !leaderboard.HasFieldName(sortKey) {
		return nil, ErrEntryInvalidSortKey
	}

	sk := sortKey
	if sk == "" {
		sk = leaderboard.FieldNames[0]
	}

	entries := []*Entry{}
	for _, entry := range leaderboard.Entries {
		entries = append(entries, entry)
	}

	slices.SortStableFunc[[]*Entry](entries, func(a, b *Entry) int {
		if order == "asc" {
			return a.FieldNamesToValues[sk] - b.FieldNamesToValues[sk]
		}
		return b.FieldNamesToValues[sk] - a.FieldNamesToValues[sk]
	})

	return entries, nil
}
