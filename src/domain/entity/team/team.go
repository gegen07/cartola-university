package team

import (
	"github.com/gegen07/cartola-university/domain/entity/player"
	"github.com/gegen07/cartola-university/domain/entity/tournament"
	"time"
)

// Team is a team which is playing the cup
type Team struct {
	ID             int                `gorm:"primary_key;auto_increment" json:"id"`
	Name           string             `gorm:"not null" json:"name"`
	Nickname       string             `gorm:"not null" json:"nickname"`
	ImageShield    string             `gorm:"size:255;null;" json:"image_shield"`
	Players        []player.Player    `gorm:"foreignkey:team_id" json:"players"`
	HomeMatches    []tournament.Match `gorm:"foreignkey:home_team_id" json:"home_matches"`
	VisitorMatches []tournament.Match `gorm:"foreignkey:visitor_team_id" json:"visitor_matches"`
	Stats          TeamStats          `gorm:"foreignkey:team_id" json:"stats"`
	CreatedAt      time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicTeam struct {
	ID             int                      `json:"id"`
	Name           string                   `json:"name"`
	Nickname       string                   `json:"nickname"`
	ImageShield    string                   `json:"image_shield"`
	Players        []player.Player          `json:"players"`
	HomeMatches    []tournament.PublicMatch `json:"home_matches"`
	VisitorMatches []tournament.PublicMatch `json:"visitor_matches"`
	Stats          PublicTeamStats          `json:"stats"`
}
