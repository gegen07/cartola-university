package persistence

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
	"github.com/jinzhu/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *PositionRepository {
	return &PositionRepository{
		db: db,
	}
}

var _ repository.PositionRepository = &PositionRepository{}

func (p PositionRepository) GetAll(args ...interface{}) ([]scout.Position, error) {
	var positions []scout.Position
	err := p.db.Debug().Preload("Scouts").Find(&positions).Error

	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (p PositionRepository) GetById(id uint64) (*scout.Position, error) {
	var position scout.Position

	err := p.db.Debug().Preload("Scouts").Where("id = ?", id).Take(&position).Error

	if err != nil {
		return nil, err
	}

	return &position, nil
}

func (p PositionRepository) Insert(position *scout.Position) (*scout.Position, error) {
	err := p.db.Debug().Create(position).Error

	if err != nil {
		return nil, err
	}

	return position, nil
}

func (p PositionRepository) AppendScoutAssociation(position *scout.Position, scout *scout.Scout) error {
	err := p.db.Debug().Model(position).Association("Scouts").Append(scout).Error

	return err
}

func (p PositionRepository) Update(position *scout.Position) (*scout.Position, error) {
	err := p.db.Debug().Save(position).Error

	if err != nil {
		return nil, err
	}

	return position, nil
}

func (p PositionRepository) Delete(id uint64) error {
	err := p.db.Debug().Where("id = ?", id).Delete(&scout.Position{}).Error

	if err != nil {
		return err
	}

	return nil
}

