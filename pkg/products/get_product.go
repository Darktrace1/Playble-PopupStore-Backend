package products

import (
	"context"

	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (h handler) GetProduct(c *fiber.Ctx) error{
	id := c.Params("cid")

	var product models.Product

	filter := bson.M{"cid": id}

	if result := h.DB.Read(context.TODO(), "new_tech", "products", &filter, &product); result != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}