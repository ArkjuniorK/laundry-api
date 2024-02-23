package laundry

import core "github.com/ArkjuniorK/laundry-api/internal/core"

func New(db *core.Database, api *core.API, lg *core.Logger, ps *core.PubSub) {
	var (
		rpo = &repository{db}
		svc = &service{rpo, lg}
		hdl = &handler{svc, rpo}
		rtr = &router{api, ps, lg, hdl}
	)

	rtr.Serve()
}
