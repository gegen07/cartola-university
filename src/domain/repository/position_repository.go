package repository

import "github.com/gegen07/cartola-university/domain/entity/scout"

type PositionRepository interface {
	GetAll(args...interface{}) ([]scout.Position, error)
	GetById(id uint64) (*scout.Position, error)
	Insert(position *scout.Position) (*scout.Position, error)
	Update(position *scout.Position) (*scout.Position, error)
	Delete(id uint64) error
	AppendScoutAssociation(position *scout.Position, scout *scout.Scout) error
}