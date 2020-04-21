package player

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/entity/team"
)

// Player struct is the data of players of Team
type Player struct {
	ID         int `gorm:"primary_key; auto_increment" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Nickname   string `gorm:"not null" json:"nickname"`
	Photo      string `gorm:"not null" json:"photo"`
	TeamID     uint `gorm:"column:team_id"`
	StatsMatches []StatsMatch
	Team       team.Team
	Position   scout.Position
	Status     Status
}
