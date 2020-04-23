package player

import "github.com/gegen07/cartola-university/domain/entity/tournament"

type StatsMatch struct {
	ID uint64  `gorm:"primary_key;auto_increment" json:"id"`
	MatchID uint64 `gorm:"foreign_key:id" json:"match_id"`
	Match 	tournament.Match `json:"match"`
	PlayerID uint64 `gorm:"column: player_id" json:"player_id"`
	Scouts []Stats `gorm:"foreign_key:stats_match_id" json:"scouts"`
}

type PublicStatsMatch struct {
	ID uint64  `gorm:"primary_key;auto_increment" json:"id"`
	PlayerID uint64 `gorm:"column: player_id" json:"player_id"`
	Match 	tournament.Match `json:"match"`
	Scouts []Stats `gorm:"foreign_key:stats_match_id" json:"scouts"`
}