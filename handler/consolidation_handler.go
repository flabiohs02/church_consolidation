package handler

import (
	"strconv"

	"church_consolidation/domain"
	"church_consolidation/usecase"

	"github.com/gofiber/fiber/v2"
)

type ConsolidationHandler struct {
	consolidationService *usecase.ConsolidationService
}

func NewConsolidationHandler(service *usecase.ConsolidationService) *ConsolidationHandler {
	return &ConsolidationHandler{consolidationService: service}
}

func (h *ConsolidationHandler) CreateConsolidation(c *fiber.Ctx) error {
	var consolidation domain.Consolidation
	if err := c.BodyParser(&consolidation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.consolidationService.CreateConsolidation(&consolidation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(consolidation)
}

func (h *ConsolidationHandler) GetConsolidationByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid consolidation ID"})
	}

	consolidation, err := h.consolidationService.GetConsolidationByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Consolidation not found"})
	}

	return c.Status(fiber.StatusOK).JSON(consolidation)
}

func (h *ConsolidationHandler) GetAllConsolidations(c *fiber.Ctx) error {
	consolidations, err := h.consolidationService.GetAllConsolidations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(consolidations)
}

func (h *ConsolidationHandler) UpdateConsolidation(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid consolidation ID"})
	}

	var consolidation domain.Consolidation
	if err := c.BodyParser(&consolidation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	consolidation.ID = uint(id)
	if err := h.consolidationService.UpdateConsolidation(&consolidation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(consolidation)
}

func (h *ConsolidationHandler) DeleteConsolidation(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid consolidation ID"})
	}

	if err := h.consolidationService.DeleteConsolidation(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
} 