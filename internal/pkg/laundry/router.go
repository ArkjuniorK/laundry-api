package laundry

import core "github.com/ArkjuniorK/laundry-api/internal/core"

type router struct {
	api     *core.API
	pubsub  *core.PubSub
	logger  *core.Logger
	handler *handler
}

func (r *router) Serve() {

	h := r.handler
	api := r.api.GetRouter().Group("/laundry")

	serviceApi := api.Group("/services")
	{
		serviceApi.Get("", h.Services)
		serviceApi.Post("", h.CreateService)
		serviceApi.Get("/:id", h.DetailService)
		serviceApi.Put("/:id", h.UpdateService)
		serviceApi.Delete("/:id", h.DeleteService)
	}

	durationApi := api.Group("/durations")
	{
		durationApi.Get("", h.Durations)
		durationApi.Post("", h.CreateDuration)
		durationApi.Get("/:id", h.DeleteDuration)
		durationApi.Put("/:id", h.UpdateDuration)
		durationApi.Delete("/:id", h.DeleteDuration)
	}

	transactionApi := api.Group("/transactions")
	{
		transactionApi.Get("", h.Transactions)
		transactionApi.Post("", h.CreateTransaction)
		transactionApi.Get("/:id", h.DetailTransaction)
		transactionApi.Put("/:id", h.UpdateTransaction)
		transactionApi.Delete("/:id", h.DeleteTransaction)
	}

}
