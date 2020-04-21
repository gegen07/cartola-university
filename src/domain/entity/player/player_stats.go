package player

// PlayerStats represents stats of a Player
type Stats struct {
	ID           uint64
	StatsMatchID uint64 `gorm:"column:stats_match_id"`
	ScoutID 	 uint64 `gorm:"column:scout_id"`
	AmountScout  uint64
}
