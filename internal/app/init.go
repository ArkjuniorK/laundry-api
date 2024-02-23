package app

import (
	"github.com/ArkjuniorK/laundry-api/internal/core"
	"github.com/ArkjuniorK/laundry-api/internal/middleware"
	"github.com/ArkjuniorK/laundry-api/internal/pkg/laundry"
	slogfiber "github.com/samber/slog-fiber"
)

// initCore initialize core packages of application
// such as router, logger, database, etc.
func (app *app) initCore() {

	logger := core.NewLogger()
	app.Logger = logger

	db := core.NewDatabase(logger)
	app.DB = db

	api := core.NewApi(logger)
	middleware.InitAPIMiddleware(api, slogfiber.New(logger.GetCore()))
	app.API = api

	// for now pubsub is not require by the system
	// pubSub := core.NewPubSub(logger)
	// middleware.InitPubSubMiddleware(pubSub)
	//app.PubSub = pubSub

	defer logger.GetCore().Info("Initialize dependencies done!")

}

// initPackages initialize all the packages inside the pkg directory.
// This function act as single source where all
// packages should be initialized.
func (app *app) initPackages() {

	var (
		db     = app.DB
		api    = app.API
		logger = app.Logger
		// pubSub = app.PubSub
	)

	laundry.New(db, api, logger, nil)

	defer logger.GetCore().Info("Initialize packages done!")
}
