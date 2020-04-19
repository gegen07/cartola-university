package entity

// Team is a team which is playing the cup
type Team struct {
	ID          int
	Name        string
	Nickname    string
	ImageShield string
	Players     []Player
	Stats       *TeamStats
}
