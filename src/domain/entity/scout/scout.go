package scout

import (
	"time"
)

// Scout struct represents the stats of each player
type Scout struct {
	ID           	uint64 `json:"id"`
	Description  	string `json:"description"`
	Points       	float64 `json:"points"`
	CreatedAt    	time.Time  `json:"created_at"`
	UpdatedAt 		time.Time  `json:"updated_at"`
}

type PublicScout struct {
	ID          uint64 `json:"id"`
	Description string `json:"description"`
	Points      float64 `json:"points"`
	PositionID	uint64 `json:"position_id"`
}

type Scouts []Scout

func (scouts Scouts) PublicScouts() []interface{} {
	result := make([]interface{}, len(scouts))

	for i, scout := range scouts {
		result[i] = scout.PublicScout()
	}

	return result
}

func (s *Scout) PublicScout() *PublicScout {
	return &PublicScout{
		ID:          s.ID,
		Description: s.Description,
		Points:      s.Points,
	}
}

func (s *Scout) Prepare() {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}
