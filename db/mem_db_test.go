package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateEntry_OnlyOverwriteSpecifiedFields(t *testing.T) {
	database := NewMemDB()
	leaderboard := database.CreateLeaderboard(
		"test leaderboard",
		[]string{"score", "rank"},
		nil,
	)
	entry, _ := database.CreateEntry(
		leaderboard.ID,
		"reez",
		map[string]int{
			"score": 10,
			"rank": 99,
		},
	)

	expected := &Entry{
		ID: entry.ID,
		Name: entry.Name,
		FieldNamesToValues: map[string]int{
			"score": 69,
			"rank": 99,
		},
	}

	// Expecting "score" to change to 69, but "rank" to stay as 99.
	actual, _ := database.UpdateEntry(
		leaderboard.ID,
		entry.ID,
		map[string]int{"score": 69},
	)

	assert.Equal(t, expected, actual)
}
