package laundry

import (
	"github.com/ArkjuniorK/laundry-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Transactions(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = c.QueryInt("limit", 10)
	)

	ts, err := h.repository.FindTransactions(c.UserContext(), page, limit)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, ts, nil)
}

func (h *handler) CreateTransaction(c *fiber.Ctx) error {
	var tr TransactionRequest

	if err := c.BodyParser(&tr); err != nil {
		return err
	}

	t, err := h.service.CreateTransaction(c.UserContext(), &tr)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, t, nil)
}

func (h *handler) DetailTransaction(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	t, err := h.repository.FindOneTransaction(c.UserContext(), id)
	if err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, t, nil)
}

func (h *handler) UpdateTransaction(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	var tr TransactionRequest
	if err := c.BodyParser(&tr); err != nil {
		return err
	}

	if err := h.service.UpdateTransaction(c.UserContext(), id, &tr); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}

func (h *handler) DeleteTransaction(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 1)
	if err != nil {
		return err
	}

	if err := h.repository.DeleteTransaction(c.UserContext(), id); err != nil {
		return err
	}

	return utils.ResponseDispatcher(c, nil, nil)
}
