package tournament

import (
	"time"
)

// Round struct represents the matches of the week
type Round struct {
	ID              uint64
	TournamentID	uint64 `gorm:"foreign_key:tournament_id" json:"tournament_id"`
	RoundBeginDate  time.Time
	RoundFinishDate time.Time
}
