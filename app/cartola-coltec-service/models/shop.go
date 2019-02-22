package models

import (
	"time"
)

// Shop struct represent the stats of shop
type Shop struct {
	CurrentRound 		Round
	Status		 		string
	Players      		[]Player
	Year         		int
	Cup                	string
	ClosingDatetime		time.Time	
}