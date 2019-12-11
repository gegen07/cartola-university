package models

// TeamStats struct represents stats of the team
type TeamStats struct {
	ID 				int
	Team      		Team
	Victory			int
	Lose			int
	Draw			int
	GoalAgainst		int
	GoalDifference	int
	GoalPro			int
	YellowCards		int
	RedCards		int
}