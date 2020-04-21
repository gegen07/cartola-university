package tournament

import (
	"github.com/gegen07/cartola-university/domain/entity/player"
)

// Shop struct represent the stats of shop
type Shop struct {
	ID           int
	Status       string
	Players      []player.Player
	Year         int
	CurrentRoundID uint64
	TournamentID uint64
}
