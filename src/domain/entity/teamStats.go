package entity

// TeamStats struct represents stats of the team
type TeamStats struct {
	ID             uint64 `gorm:"primary_key;auto_increment" json:"id"`
	TeamRefer      uint64 `gorm:"not null" json:"TeamRefer"`
	Victory        uint64 `gorm:"not null" json:"victory"`
	Lose           uint64 `gorm:"not null" json:"lose"`
	Draw           uint64 `gorm:"not null" json:"draw"`
	GoalAgainst    uint64 `gorm:"not null" json:"goal_against"`
	GoalDifference uint64 `gorm:"not null" json:"goal_difference"`
	GoalPro        uint64 `gorm:"not null" json:"goal_pro"`
	YellowCards    uint64 `gorm:"not null" json:"yellow_cards"`
	RedCards       uint64 `gorm:"not null" json:"red_cards"`
}
