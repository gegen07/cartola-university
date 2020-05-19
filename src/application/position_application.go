package application

import (
	"context"
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
	"golang.org/x/sync/errgroup"
	"time"
)

type PositionApplicationInterface interface {
	GetAll(ctx context.Context, page int) ([]scout.Position, error)
	GetById(ctx context.Context, id uint64) (*scout.Position, error)
	Insert(ctx context.Context, position *scout.Position) (*scout.Position, error)
	Update(ctx context.Context, position *scout.Position) (*scout.Position, error)
	Delete(ctx context.Context, id uint64) error
}

type PositionApplication struct {
	repo                    repository.PositionRepository
	scoutPositionRepository repository.ScoutPositionRepository
	contextTimeout time.Duration
}

var _ PositionApplicationInterface = &PositionApplication{}

func (p *PositionApplication) fillPositionWithScout(c context.Context, data []scout.Position) ([]scout.Position, error) {
	g, ctx := errgroup.WithContext(c)

	for _, position := range data  {
		g.Go(func() error {
			res, err := p.scoutPositionRepository.GetScoutsByPositionID(ctx, position.ID)

			if err != nil {
				return err
			}

			position.Scouts = res

			return nil
		})
	}

	go func() {
		err := g.Wait()

		if err != nil {
			return
		}
	}()

	if err := g.Wait(); err!= nil {
		return nil, err
	}

	return data, nil
}

func (p *PositionApplication) GetAll(ctx context.Context, page int) ([]scout.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	positions, err := p.repo.GetAll(ctx, page)

	if err != nil {
		return nil, err
	}

	return p.fillPositionWithScout(ctx, positions)
}

func (p *PositionApplication) GetById(ctx context.Context,id uint64) (*scout.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	position, err := p.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	res, err := p.scoutPositionRepository.GetScoutsByPositionID(ctx, position.ID)

	if err != nil {
		return nil, err
	}

	position.Scouts = res

	return position, nil
}

func (p *PositionApplication) Insert(ctx context.Context, position *scout.Position) (*scout.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.repo.Insert(ctx, position)
}

func (p *PositionApplication) Update(ctx context.Context, position *scout.Position) (*scout.Position, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.repo.Update(ctx, position)
}

func (p *PositionApplication) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.repo.Delete(ctx, id)
}