package repository

import (
	"github.com/gegen07/cartola-university/domain/entity"
)

// Repository represent the round's repository contract
type FormationRepository interface {
	GetAll(args ...interface{}) ([]entity.Formation, error)
	GetByID(id uint64) (*entity.Formation, error)
	Insert(formation *entity.Formation) (*entity.Formation, error)
	Update(formation *entity.Formation) (*entity.Formation, error)
	Delete(id uint64) error
}
