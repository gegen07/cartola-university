package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Formation is the schema of game team
type Formation struct {
	gorm.Model
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Goalkeeper uint64 `gorm:"not null" json:"goalkeeper"`
	Attackers  uint64 `gorm:"not null" json:"attackers"`
	Defenders  uint64 `gorm:"not null" json: "defenders"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PublicFormation struct {
	gorm.Model
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Goalkeeper uint64 `gorm:"not null" json:"goalkeeper"`
	Attackers  uint64 `gorm:"not null" json:"attackers"`
	Defenders  uint64 `gorm:"not null" json: "defenders"`
}

type Formations []Formation

func (formations Formations) PublicFormations() []interface{} {
	result := make([]interface{}, len(formations))

	for index, formation := range formations {
		result[index] = formation.PublicFormation()
	}

	return result
}

func (f *Formation) PublicFormation() interface{} {
	return &PublicFormation {
		ID: f.ID,
		Goalkeeper: f.Goalkeeper,
		Attackers: f.Attackers,
		Defenders: f.Defenders,
	}
}

func (f *Formation) Prepare() {
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}


