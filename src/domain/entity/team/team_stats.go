package team

import "time"

// TeamStats struct represents stats of the team
type TeamStats struct {
	ID             	uint64 `gorm:"primary_key;auto_increment" json:"id"`
	TeamID         	uint64 `gorm:"not null; column:team_id" json:"TeamID"`
	Victory        	uint64 `gorm:"not null" json:"victory"`
	Lose           	uint64 `gorm:"not null" json:"lose"`
	Draw           	uint64 `gorm:"not null" json:"draw"`
	GoalAgainst    	uint64 `gorm:"not null" json:"goal_against"`
	GoalDifference 	uint64 `gorm:"not null" json:"goal_difference"`
	GoalPro        	uint64 `gorm:"not null" json:"goal_pro"`
	YellowCards    	uint64 `gorm:"not null" json:"yellow_cards"`
	RedCards       	uint64 `gorm:"not null" json:"red_cards"`
	CreatedAt 	   	time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	   	time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicTeamStats struct {
	ID             	uint64 `json:"id"`
	TeamID        	uint64 `json:"team_id"`
	Victory        	uint64 `json:"victory"`
	Lose           	uint64 `json:"lose"`
	Draw           	uint64 `json:"draw"`
	GoalAgainst    	uint64 `json:"goal_against"`
	GoalDifference 	uint64 `json:"goal_difference"`
	GoalPro        	uint64 `json:"goal_pro"`
	YellowCards    	uint64 `json:"yellow_cards"`
	RedCards       	uint64 `json:"red_cards"`
}
