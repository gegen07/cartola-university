package entity

// PlayerStats represents stats of a Player
type PlayerStats struct {
	ID           int
	AmountScout  int
	CurrentRound Round
	Player       Player
	Scout        Scout
}
