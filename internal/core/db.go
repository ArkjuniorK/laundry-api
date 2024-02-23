package core

import (
	"os"
	"path"

	"github.com/ArkjuniorK/laundry-api/internal/model"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	core *gorm.DB
}

func NewDatabase(l *Logger) *Database {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logger := slogGorm.New(slogGorm.WithLogger(l.GetCore()), slogGorm.WithTraceAll())
	option := &gorm.Config{Logger: logger}

	dialect := sqlite.Open(path.Join(wd, "laundry.db"))
	db, err := gorm.Open(dialect, option)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Service{}, &model.Transaction{}, &model.Duration{})

	return &Database{core: db}
}

func (db *Database) GetCore() *gorm.DB {
	return db.core
}
