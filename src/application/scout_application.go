package application

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
)

type ScoutApplication struct {
	repo repository.ScoutRepository
}

type ScoutApplicationInterface interface {
	GetAll(args ...interface{}) ([]scout.Scout, error)
	GetByID(id uint64) (*scout.Scout, error)
	Insert(scout *scout.Scout) (*scout.Scout, error)
	Update(scout *scout.Scout) (*scout.Scout, error)
	Delete(id uint64) error
}

var _ ScoutApplicationInterface = &ScoutApplication{}

func (s ScoutApplication) GetAll(args ...interface{}) ([]scout.Scout, error) {
	return s.repo.GetAll(args)
}

func (s ScoutApplication) GetByID(id uint64) (*scout.Scout, error) {
	return s.repo.GetByID(id)
}

func (s ScoutApplication) Insert(scout *scout.Scout) (*scout.Scout, error) {
	return s.repo.Insert(scout)
}

func (s ScoutApplication) Update(scout *scout.Scout) (*scout.Scout, error) {
	return s.repo.Update(scout)
}

func (s ScoutApplication) Delete(id uint64) error {
	return s.repo.Delete(id)
}
