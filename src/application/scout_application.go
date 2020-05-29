package application

import (
	"context"
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/domain/repository"
	"golang.org/x/sync/errgroup"
	"time"
)

type ScoutApplication struct {
	repo repository.ScoutRepository
	scoutPositionRepo repository.ScoutPositionRepository
	contextTimeout time.Duration
}

type ScoutApplicationInterface interface {
	GetAll(ctx context.Context, page int) ([]scout.Scout, error)
	GetByID(ctx context.Context, id uint64) (*scout.Scout, error)
	Insert(ctx context.Context, scout *scout.Scout) (*scout.Scout, error)
	Update(ctx context.Context, scout *scout.Scout) (*scout.Scout, error)
	Delete(ctx context.Context, id uint64) error
}

var _ ScoutApplicationInterface = &ScoutApplication{}

func (s *ScoutApplication) fillScoutWithPosition(c context.Context, data []scout.Scout) ([]scout.Scout, error) {
	g, ctx := errgroup.WithContext(c)

	for _, scout := range data  {
		g.Go(func() error {
			res, err := s.scoutPositionRepo.GetPositionsByScoutID(ctx, scout.ID)

			if err != nil {
				return err
			}

			scout.Positions = res

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

func (s ScoutApplication) GetAll(ctx context.Context, page int) ([]scout.Scout, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	scouts, err := s.repo.GetAll(ctx, page)

	if err != nil {
		return nil, err
	}

	return s.fillScoutWithPosition(ctx, scouts)
}

func (s ScoutApplication) GetByID(ctx context.Context, id uint64) (*scout.Scout, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	scout, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	positions, err := s.scoutPositionRepo.GetPositionsByScoutID(ctx, scout.ID)

	if err != nil {
		return nil, err
	}

	scout.Positions = positions

	return scout, nil
}

func (s ScoutApplication) Insert(c context.Context, scout *scout.Scout) (*scout.Scout, error) {
	g, ctx := errgroup.WithContext(c)

	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	g.Go(func() error {
		for _, position := range scout.Positions {
			err := s.scoutPositionRepo.AddRelation(c, scout.ID, position.ID)

			if err != nil {
				return err
			}
		}

		return nil
	})

	go func() {
		err := g.Wait()

		if err != nil {
			return
		}
	}()

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return s.repo.Insert(ctx, scout)
}

func (s ScoutApplication) Update(ctx context.Context, scout *scout.Scout) (*scout.Scout, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	return s.repo.Update(ctx, scout)
}

func (s ScoutApplication) Delete(ctx context.Context, id uint64) error {
	g, ctx := errgroup.WithContext(ctx)

	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	res, err := s.scoutPositionRepo.GetPositionsByScoutID(ctx, id)

	if err != nil {
		return err
	}

	g.Go(func() error {
		for _, r := range res {
			err := s.scoutPositionRepo.DeleteRelation(ctx, r.ID, id)

			if err != nil {
				return err
			}
		}

		return nil
	})

	go func() {
		err := g.Wait()

		if err != nil {
			return
		}
	}()

	if err := g.Wait(); err!= nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
