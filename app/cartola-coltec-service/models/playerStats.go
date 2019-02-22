package models

// PlayerStats represents stats of a Player
type PlayerStats struct {
	ID 				int
	CurrentRound	Round
	Player			Player
	AmountScout		int
	Scout			Scout		
}	