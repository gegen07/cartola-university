package round

import (
	"context"
	"github.com/gegen07/cartola-coltec/app/cartola-coltec-service/models"
)

// Repository represent the round's repository contract
type Repository interface {
	GetAll(ctx context.Context, args ...interface{}) ([]*models.Round, error)
	GetByID(ctx context.Context, id int64) (*models.Round, error)
	Insert(ctx context.Context, round *models.Round) error
	Update(ctx context.Context, round *models.Round) error
	Delete(ctx context.Context, id int64) error
}