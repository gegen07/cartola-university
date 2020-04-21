package tournament

import "time"

type Tournament struct {
	Id 			uint64
	Name 		uint64
	Rounds 		[]Round
	CreatedAt 	time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
