package persistence

import (
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
	"github.com/jinzhu/gorm"
)

type ScoutRepository struct {
	db *gorm.DB
}

func NewScoutRepository(db *gorm.DB) *ScoutRepository {
	return &ScoutRepository{
		db: db,
	}
}

var _ repository.ScoutRepository = &ScoutRepository{}

func (s ScoutRepository) GetAll(args ...interface{}) ([]scout.Scout, error) {
	var scouts []scout.Scout

	err := s.db.Debug().Find(&scouts).Error

	if err != nil {
		return nil, err
	}

	return scouts, nil
}

func (s ScoutRepository) GetByID(id uint64) (*scout.Scout, error) {
	var scout scout.Scout

	err := s.db.Debug().Where("id = ?", id).Take(&scout).Error

	if err != nil {
		return nil, err
	}

	return &scout, nil
}

func (s ScoutRepository) Insert(scout *scout.Scout) (*scout.Scout, error) {
	err := s.db.Debug().Create(scout).Error

	if err != nil {
		return nil, err
	}

	return scout, nil
}

func (s ScoutRepository) Update(scout *scout.Scout) (*scout.Scout, error) {
	err := s.db.Debug().Save(scout).Error

	if err != nil {
		return nil, err
	}

	return scout, nil
}

func (s ScoutRepository) Delete(id uint64) error {
	err := s.db.Debug().Where("id = ?", id).Delete(&scout.Scout{}).Error

	if err != nil {
		return err
	}

	return nil
}