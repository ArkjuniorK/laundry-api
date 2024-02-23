package laundry

import (
	"context"

	"github.com/ArkjuniorK/laundry-api/internal/model"
	"github.com/ArkjuniorK/laundry-api/internal/utils"
)

func (r *repository) InsertServices(ctx context.Context, svcs ...model.Service) error {
	return r.db.GetCore().Model(&model.Service{}).CreateInBatches(svcs, len(svcs)).Error
}

func (r *repository) FindServices(ctx context.Context, page, limit int) ([]model.Service, error) {
	var (
		db   = r.db.GetCore()
		svcs = make([]model.Service, 0)
	)

	if err := db.Model(&model.Service{}).Scopes(utils.GormPaginator(page, limit)).Find(&svcs).Error; err != nil {
		return nil, err
	}

	return svcs, nil
}

func (r *repository) FindOneService(ctx context.Context, id int) (*model.Service, error) {
	var (
		db  = r.db.GetCore()
		svc = new(model.Service)
	)

	if err := db.Model(&model.Service{}).First(svc, id).Error; err != nil {
		return nil, err
	}

	return svc, nil
}

func (r *repository) UpdateService(ctx context.Context, id int, svc *model.Service) error {
	return r.db.GetCore().Model(&model.Service{}).Where("id =?", id).Updates(svc).Error
}

func (r *repository) DeleteService(ctx context.Context, id int) error {
	return r.db.GetCore().Delete(&model.Service{}, "id = ?", id).Error
}
