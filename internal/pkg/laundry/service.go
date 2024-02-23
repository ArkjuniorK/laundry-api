package laundry

import (
	"context"
	"strconv"
	"time"

	core "github.com/ArkjuniorK/laundry-api/internal/core"
	"github.com/ArkjuniorK/laundry-api/internal/model"
)

type service struct {
	repo   *repository
	logger *core.Logger
}

func (s *service) CreateTransaction(ctx context.Context, pyd *TransactionRequest) (*model.Transaction, error) {
	t, err := s.setTransaction(ctx, pyd)
	if err != nil {
		return nil, err
	}

	if err = s.repo.InsertTransaction(ctx, *t); err != nil {
		return nil, err
	}

	return t, nil
}

func (s *service) UpdateTransaction(ctx context.Context, id int, pyd *TransactionRequest) error {
	t, err := s.setTransaction(ctx, pyd)
	if err != nil {
		return err
	}

	if err = s.repo.UpdateTransaction(ctx, id, t); err != nil {
		return err
	}

	return nil
}

func (s *service) setTransaction(ctx context.Context, pyd *TransactionRequest) (*model.Transaction, error) {
	var t = new(model.Transaction)
	t.CustomerName = pyd.CustomerName
	t.CustomerPhone = pyd.CustomerPhone

	svc, err := s.repo.FindOneService(ctx, int(pyd.ServiceID))
	if err != nil {
		return nil, err
	}

	t.Service = *svc

	price, err := strconv.Atoi(svc.Price)
	if err != nil {
		return nil, err
	}

	var (
		d                      = 0 * time.Second
		dur                    = new(model.Duration)
		additionalCost float64 = 0
	)

	if pyd.DurationID != 0 {
		dur, err = s.repo.FindOneDuration(ctx, int(pyd.DurationID))
		if err != nil {
			return nil, err
		}

		additionalCost = (float64(dur.Cost) / 100) * float64(price)

		d, err = time.ParseDuration(dur.Duration)
		if err != nil {
			return nil, err
		}

		t.Duration = *dur
	}

	t.TotalPrice = int(float64(price) + additionalCost)
	t.TakeOn = time.Now().Add(d)

	return t, nil
}
