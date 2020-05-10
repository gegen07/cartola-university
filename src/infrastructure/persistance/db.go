package persistence

import (
	"fmt"
	"github.com/gegen07/cartola-university/domain/entity"
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	Formation repository.FormationRepository
	Scout repository.ScoutRepository
	Position repository.PositionRepository
	db *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPass, DbPort, DbHost, DbName string) (*Repositories, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPass)
	db, err := gorm.Open(DbDriver, url)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &Repositories{
		Formation: NewFormationRepository(db),
		Scout: NewScoutRepository(db),
		Position: NewPositionRepository(db),
		db: db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}

func (s *Repositories) Migrate() error {
	return s.db.AutoMigrate(
		&entity.Formation{},
		&scout.Scout{},
		&scout.Position{},
		).Error
}