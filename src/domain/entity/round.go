package entity

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Round struct represents the matches of the week
type Round struct {
	ID              int
	Matches         []Match
	RoundBeginDate  time.Time
	RoundFinishDate time.Time
}
