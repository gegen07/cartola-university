package persistence

import (
	"github.com/gegen07/cartola-university/domain/entity"
	"github.com/gegen07/cartola-university/domain/repository"
	"github.com/jinzhu/gorm"
)

type FormationRepository struct {
	db *gorm.DB
}

func NewFormationRepository(db *gorm.DB) *FormationRepository {
	return &FormationRepository{db: db}
}

var _ repository.FormationRepository= &FormationRepository{}

func (r FormationRepository) GetAll(args ...interface{}) ([]entity.Formation, error) {
	var formations []entity.Formation

	err := r.db.Debug().Find(&formations).Error

	if err != nil {
		return nil, err
	}

	return formations, nil
}

func (r FormationRepository) GetByID(id uint64) (*entity.Formation, error) {
	var formation entity.Formation

	err := r.db.Debug().Where("id = ?").Take(&formation).Error

	if err != nil {
		return nil, err
	}

	return &formation, nil
}

func (r FormationRepository) Insert(formation *entity.Formation) (*entity.Formation, error) {
	err := r.db.Debug().Create(&formation).Error

	if err != nil {
		return nil, err
	}

	return formation, nil
}

func (r FormationRepository) Update(formation *entity.Formation) (*entity.Formation, error) {
	err := r.db.Debug().Save(&formation).Error

	if err != nil {
		return nil, err
	}

	return formation, nil
}

func (r FormationRepository) Delete(formation *entity.Formation) error {
	err := r.db.Debug().Delete(&formation).Error

	if err != nil {
		return err
	}

	return nil
}