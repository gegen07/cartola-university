package entity

// Team is a team which is playing the cup
type Team struct {
	ID          		int `gorm:"primary_key;auto_increment" json:"id"`
	Name        		string `gorm:"not null" json:"name"`
	Nickname    		string `gorm:"not null" json:"nickname"`
	ImageShield 		string `gorm:"size:255;null;" json:"image_shield"`
	Refer 				uint `json:"-"`
	Players     		[]Player `gorm:"foreignkey:TeamRefer association_foreignkey:Refer" json:"players"`
	HomeMatches 		[]Match `gorm:"foreignkey:HomeTeam association_foreignkey:Refer" json:"home_matches"`
	VisitorMatches 		[]Match `gorm:"foreignkey:VisitorTeam association_foreignkey:Refer" json:"visitor_matches"`
	Stats				TeamStats `gorm:"foreignkey:TeamRefer association_foreignkey:Refer" json:"stats"`
}
