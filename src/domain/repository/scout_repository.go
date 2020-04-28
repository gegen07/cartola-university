package repository

import "github.com/gegen07/cartola-university/domain/entity/scout"

type ScoutRepository interface {
	GetAll(args ...interface{}) ([]scout.Scout, error)
	GetByID(id uint64) (*scout.Scout, error)
	Insert(scout *scout.Scout) (*scout.Scout, error)
	Update(scout *scout.Scout) (*scout.Scout, error)
	Delete(id uint64) error
}
