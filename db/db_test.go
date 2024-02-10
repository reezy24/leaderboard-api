package db

import (
	"testing"
	"reflect"
)

func TestNewLeaderboard(t *testing.T) {
	name := "Example Leaderboard"
	columns := []string{"Column1", "Column2"}

	leaderboard := NewLeaderboard(name, columns)

	expectedLeaderboard := &Leaderboard{
		ID:      leaderboard.ID,
		Name:    name,
		Columns: columns,
	}

	if !reflect.DeepEqual(leaderboard, expectedLeaderboard) {
		t.Errorf("Expected %v, but got %v", expectedLeaderboard, leaderboard)
	}
}
