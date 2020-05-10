package application

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
)

type PositionApplicationInterface interface {
	GetAll(args...interface{}) ([]scout.Position, error)
	GetById(id uint64) (*scout.Position, error)
	Insert(position *scout.Position) (*scout.Position, error)
	Update(position *scout.Position) (*scout.Position, error)
	Delete(id uint64) error
	AppendScoutAssociation(position *scout.Position, scout *scout.Scout) error
}

type PositionApplication struct {
	repo repository.PositionRepository
}

func (p *PositionApplication) AppendScoutAssociation(position *scout.Position, scout *scout.Scout) error {
	return p.repo.AppendScoutAssociation(position, scout)
}

func (p *PositionApplication) GetAll(args ...interface{}) ([]scout.Position, error) {
	return p.repo.GetAll(args)
}

func (p *PositionApplication) GetById(id uint64) (*scout.Position, error) {
	return p.repo.GetById(id)
}

func (p *PositionApplication) Insert(position *scout.Position) (*scout.Position, error) {
	return p.repo.Insert(position)
}

func (p *PositionApplication) Update(position *scout.Position) (*scout.Position, error) {
	return p.repo.Update(position)
}

func (p *PositionApplication) Delete(id uint64) error {
	return p.repo.Delete(id)
}

var _ PositionApplicationInterface = &PositionApplication{}
