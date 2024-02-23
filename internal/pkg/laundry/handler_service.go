package laundry

import (
	"github.com/ArkjuniorK/laundry-api/internal/model"
	"github.com/ArkjuniorK/laundry-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Services(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = c.QueryInt("limit", 10)
	)

	svcs, err := h.repository.FindServices(c.UserContext(), page, limit)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, svcs, nil)
}

func (h *handler) CreateService(c *fiber.Ctx) error {
	var svc model.Service

	if err := c.BodyParser(&svc); err != nil {
		return err
	}

	if err := h.repository.InsertServices(c.UserContext(), svc); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}

func (h *handler) DetailService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	svc, err := h.repository.FindOneService(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, svc, nil)
}

func (h *handler) UpdateService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	svc := new(model.Service)
	if err := c.BodyParser(svc); err != nil {
		return err
	}

	if err := h.repository.UpdateService(c.UserContext(), id, svc); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}

func (h *handler) DeleteService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	if err := h.repository.DeleteService(c.UserContext(), id); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}
