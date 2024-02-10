package db

import (
	"reflect"
	"testing"
)

func TestNewLeaderboard(t *testing.T) {
	name := "Example Leaderboard"
	columns := []string{"Column1", "Column2"}

	leaderboard := NewLeaderboard(name, columns)

	expectedLeaderboard := &Leaderboard{
		ID:         leaderboard.ID,
		Name:       name,
		FieldNames: columns,
	}

	if !reflect.DeepEqual(leaderboard, expectedLeaderboard) {
		t.Errorf("Expected %v, but got %v", expectedLeaderboard, leaderboard)
	}
}
