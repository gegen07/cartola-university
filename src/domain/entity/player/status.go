package player

import "time"

// Status struct represents
// Problably-Injured-Doubt-Suspended-Nothing
type Status struct {
	ID          int `gorm:"primary_key;auto_increment" json:"id"`
	Description string `gorm:"size:100;not null;" json:"description"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicStatus struct {
	ID          int `gorm:"primary_key;auto_increment" json:"id"`
	Description string `gorm:"size:100;not null;" json:"description"`
}

type StatusArr []Status

func (s StatusArr) PublicStatus() []interface{} {
	result := make([]interface{}, len(s))

	for i, status := range s {
		result[i] = status
	}

	return result
}

func (s *Status) PublicStatus() *PublicStatus {
	return &PublicStatus{
		ID:          s.ID,
		Description: s.Description,
	}
}

func (s *Status) Prepare() {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}