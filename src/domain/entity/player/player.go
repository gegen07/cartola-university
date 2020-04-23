package player

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/entity/team"
	"time"
)

// Player struct is the data of players of Team
type Player struct {
	ID         uint64 `gorm:"primary_key; auto_increment" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Nickname   string `gorm:"not null" json:"nickname"`
	Photo      string `gorm:"not null" json:"photo"`
	TeamID     uint64 `gorm:"column:team_id"`
	PositionID uint64 `gorm:"foreignkey:id"`
	Team 	   team.Team `json:"team"`
	Position   scout.Position `json:"position"`
	StatsMatches []StatsMatch `gorm:"foreignkey:player_id" json:"matches_stats"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicPlayer struct {
	ID         uint64 `gorm:"primary_key; auto_increment" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Nickname   string `gorm:"not null" json:"nickname"`
	Photo      string `gorm:"not null" json:"photo"`
	Team 	   team.Team `json:"team"`
	Position   scout.Position `json:"position"`
	StatsMatches []StatsMatch `gorm:"foreignkey:player_id" json:"matches_stats"`
}
