package entity

// Shop struct represent the stats of shop
type Shop struct {
	ID           int
	CurrentRound Round
	Status       string
	Players      []Player
	Year         int
	Cup          string
}
