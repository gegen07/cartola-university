package tournament

import (
	"github.com/gegen07/cartola-university/domain/entity/team"
	"time"
)

// Match represents a draw between two teams
type Match struct {
	ID                      int `gorm:"primary_key;auto_increment" json:"id"`
	HomeTeamID              uint64 `gorm:"column:home_team_id" json:"home_team_id"`
	VisitorTeamID           uint64 `gorm:"column:visitor_team_id" json:"visitor_team_id"`
	RoundID 		    	uint64 `gorm:"column:round_id" json:"round_id"`
	ScoreboardHomeTeam      int `gorm:"not null" json:"scoreboard_home_team"`
	ScoreboardVisitorTeam   int	`gorm:"not null" json:"scoreboard_visitor_team"`
	Date                    time.Time `gorm:"not null" json:"date"`
	CreatedAt 				time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 				time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicMatch struct {
	ID                    int       `json:"id"`
	ScoreboardHomeTeam    int       `json:"scoreboard_home_team"`
	ScoreboardVisitorTeam int       `json:"scoreboard_visitor_team"`
	Date                  time.Time `json:"date"`
	HomeTeam              team.Team `json:"home_team"`
	VisitorTeam           team.Team `json:"visitor_team"`
	Tournament 			  Tournament
}
