package scout

import "time"

// Position struct represents the position of the player in a team
type Position struct {
	ID         	uint64 `gorm:"primary_key;auto_increment; column:id" json:"id"`
	Description string `gorm:"not null" json:"description"`
	Scouts 		[]Scout `gorm:"foreignkey:PositionID" json:"scouts"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicPosition struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Description string `gorm:"not null" json:"description"`
	Scouts 		[]Scout `gorm:"foreignkey:position_id" json:"scouts"`
}

type Positions []Position

func (p Positions) PublicPositions() []interface{} {
	result := make([]interface{}, len(p))

	for i, position := range p {
		result[i] = position.PublicPosition()
	}

	return result
}

func (p *Position) PublicPosition() *PublicPosition {
	return &PublicPosition{
		ID: p.ID,
		Description: p.Description,
		Scouts: p.Scouts,
	}
}

func (p *Position) Prepare() {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
