package entity

import (
	"time"
)

// Match represents a draw between two teams
type Match struct {
	ID                    int
	HomeTeam              uint64
	VisitorTeam           uint64
	RoundReference 		  uint64
	ScoreboardHomeTeam    int
	ScoreboardVisitorTeam int
	Date                  time.Time
}
