package application

import (
	"github.com/gegen07/cartola-university/domain/entity"
	"github.com/gegen07/cartola-university/domain/repository"
)

type formationApplication struct {
	repo repository.FormationRepository
}

var _ FormationApplicationInterface = &formationApplication{}

type FormationApplicationInterface interface {
	GetAll(args ...interface{}) ([]entity.Formation, error)
	GetByID(id uint64) (*entity.Formation, error)
	Insert(formation *entity.Formation) (*entity.Formation, error)
	Update(formation *entity.Formation) (*entity.Formation, error)
	Delete(formationId uint64) error
}

func (f formationApplication) GetAll(args ...interface{}) ([]entity.Formation, error) {
	return f.repo.GetAll(args)
}

func (f formationApplication) GetByID(id uint64) (*entity.Formation, error) {
	return f.repo.GetByID(id)
}

func (f formationApplication) Insert(formation *entity.Formation) (*entity.Formation, error) {
	return f.repo.Insert(formation)
}

func (f formationApplication) Update(formation *entity.Formation) (*entity.Formation, error) {
	return f.repo.Update(formation)
}

func (f formationApplication) Delete(formationId uint64) error {
	return f.repo.Delete(formationId)
}
