package scout

import (
	"github.com/shopspring/decimal"
	"time"
)

// Scout struct represents the stats of each player
type Scout struct {
	ID           	uint64 `gorm:"primary_key;auto_increment" json:"id"`
	PositionId		uint64 `gorm:"column:position_id" json:"position_id"`
	Description  	string `gorm:"not null" json:"description"`
	Points       	decimal.Decimal `gorm:"not null" json:"points"`
	CreatedAt    	time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicScout struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Description string `gorm:"not null" json:"description"`
	Points      decimal.Decimal `gorm:"not null" json:"points"`
	PositionId	uint64 `gorm:"column:position_id" json:"position_id"`
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
		PositionId:  s.PositionId,
	}
}

func (s *Scout) Prepare() {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}
