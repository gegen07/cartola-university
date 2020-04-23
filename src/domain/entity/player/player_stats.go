package player

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"time"
)

// PlayerStats represents stats of a Player
type Stats struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	StatsMatchID uint64 `gorm:"column:stats_match_id"`
	ScoutID 	 uint64 `gorm:"foreignkey:id; column:scout_id"`
	Scout        scout.Scout `json:"scout"`
	AmountScout  uint64 `gorm:"not null" json:"amount_scout"`
	CreatedAt    	time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 		time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicStats struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	StatsMatchID uint64 `gorm:"column:stats_match_id"`
	Scout        scout.Scout `json:"scout"`
	AmountScout  uint64 `gorm:"not null" json:"amount_scout"`
}
