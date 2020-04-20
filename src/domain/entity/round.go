package entity

import (
	"time"
)

// Round struct represents the matches of the week
type Round struct {
	ID              int
	Matches         []Match `gorm:"foreignkey:RoundReference"`
	RoundBeginDate  time.Time
	RoundFinishDate time.Time
}
