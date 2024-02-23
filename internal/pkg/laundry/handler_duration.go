package laundry

import (
	"github.com/ArkjuniorK/laundry-api/internal/model"
	"github.com/ArkjuniorK/laundry-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Durations(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = c.QueryInt("limit", 10)
	)

	durs, err := h.repository.FindSDuration(c.UserContext(), page, limit)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, durs, nil)
}

func (h *handler) CreateDuration(c *fiber.Ctx) error {
	var dur model.Duration

	if err := c.BodyParser(&dur); err != nil {
		return err
	}

	if err := h.repository.InsertDuration(c.UserContext(), dur); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}

func (h *handler) DetailDuration(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	dur, err := h.repository.FindOneDuration(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, dur, nil)
}

func (h *handler) UpdateDuration(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	dur := new(model.Duration)
	if err := c.BodyParser(dur); err != nil {
		return err
	}

	if err := h.repository.UpdateDuration(c.UserContext(), id, dur); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}

func (h *handler) DeleteDuration(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	if err := h.repository.DeleteDuration(c.UserContext(), id); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}
