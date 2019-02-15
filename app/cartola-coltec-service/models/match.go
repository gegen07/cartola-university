package models

import (
	"time"
)

// Match represents a draw between two teams
type Match struct {
	ID 						int
	HomeTeam 				Team
	VisitorTeam				Team
	ScoreboardHomeTeam 		int
	ScoreboardVisitorTeam 	int
	Date 	 				time.Time
}