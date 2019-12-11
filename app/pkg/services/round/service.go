package round

import (
	"context"
	"github.com/gegen07/cartola-coltec/app/pkg/models"
	"github.com/gegen07/cartola-coltec/app/pkg/repository/round"
)

// Service struct which is UseCase of Round Repository
type Service struct {
	repository round.Repository
}

// NewService func return a service struct with repository
func (s *Service) NewService(r round.Repository) *Service {
	return &Service {
		repository: r,
	}
}

//GetByID a instance of round
func (s *Service) GetByID(ctx context.Context, id int64) (*models.Round, error) {
	return s.repository.GetByID(ctx, id)
}

//GetAll instances of round from repository
func (s *Service) GetAll(ctx context.Context, args ...interface{}) ([]*models.Round, error) {
	return s.repository.GetAll(ctx, args)
}

//Create func insert in repository a new Round
func (s *Service) Create(ctx context.Context, round *models.Round) (error) {
	err := s.repository.Insert(ctx, round)
	return err
}

//Update func send alterations to repository
func (s *Service) Update(ctx context.Context, round *models.Round) (error) {
	return s.repository.Update(ctx, round)
}

//Delete func delete one instance of round model
func (s *Service) Delete(ctx context.Context, id int64) (error) {
	_, err := s.GetByID(ctx, id)

	if err != nil {
		return s.repository.Delete(ctx, id)
	}

	return err
}