package products

import (
	"context"

	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.ReadAll(context.TODO(), "new_tech", "products", &products); result != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}