package laundry

import (
	"context"

	"github.com/ArkjuniorK/laundry-api/internal/model"
	"github.com/ArkjuniorK/laundry-api/internal/utils"
	"gorm.io/gorm/clause"
)

func (r *repository) InsertTransaction(ctx context.Context, ts ...model.Transaction) error {
	return r.db.GetCore().Model(&model.Transaction{}).CreateInBatches(ts, len(ts)).Error
}

func (r *repository) FindTransactions(ctx context.Context, page, limit int) ([]model.Transaction, error) {
	var (
		db = r.db.GetCore()
		ts = make([]model.Transaction, 0)
	)

	if err := db.Model(&model.Transaction{}).Scopes(utils.GormPaginator(page, limit)).Preload(clause.Associations).Find(&ts).Error; err != nil {
		return nil, err
	}

	return ts, nil
}

func (r *repository) FindOneTransaction(ctx context.Context, id int) (*model.Transaction, error) {
	var (
		db = r.db.GetCore()
		t  = new(model.Transaction)
	)

	if err := db.Model(&model.Transaction{}).Preload(clause.Associations).First(t, id).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (r *repository) UpdateTransaction(ctx context.Context, id int, t *model.Transaction) error {
	return r.db.GetCore().Model(&model.Transaction{}).Where("id =?", id).Updates(t).Error
}

func (r *repository) DeleteTransaction(ctx context.Context, id int) error {
	return r.db.GetCore().Delete(&model.Transaction{}, "id = ?", id).Error
}
