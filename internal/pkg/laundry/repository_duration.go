package laundry

import (
	"context"

	"github.com/ArkjuniorK/laundry-api/internal/model"
	"github.com/ArkjuniorK/laundry-api/internal/utils"
)

func (r *repository) InsertDuration(ctx context.Context, durs ...model.Duration) error {
	return r.db.GetCore().Model(&model.Duration{}).CreateInBatches(durs, len(durs)).Error
}

func (r *repository) FindSDuration(ctx context.Context, page, limit int) ([]model.Duration, error) {
	var (
		db   = r.db.GetCore()
		durs = make([]model.Duration, 0)
	)

	if err := db.Model(&model.Duration{}).Scopes(utils.GormPaginator(page, limit)).Find(&durs).Error; err != nil {
		return nil, err
	}

	return durs, nil
}

func (r *repository) FindOneDuration(ctx context.Context, id int) (*model.Duration, error) {
	var (
		db  = r.db.GetCore()
		dur = new(model.Duration)
	)

	if err := db.Model(&model.Duration{}).First(dur, id).Error; err != nil {
		return nil, err
	}

	return dur, nil
}

func (r *repository) UpdateDuration(ctx context.Context, id int, dur *model.Duration) error {
	return r.db.GetCore().Model(&model.Duration{}).Where("id =?", id).Updates(dur).Error
}

func (r *repository) DeleteDuration(ctx context.Context, id int) error {
	return r.db.GetCore().Delete(&model.Duration{}, "id = ?", id).Error
}
