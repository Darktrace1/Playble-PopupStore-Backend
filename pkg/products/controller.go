package products

import (
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/mongorm"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	DB 	*mongorm.Model
	S3 	*manager.Uploader
}

func RegisterRoutes(app *fiber.App, db *mongorm.Model, s3 *manager.Uploader) {
    h := &handler{
        DB: db,
		S3: s3,
    }

    routes := app.Group("/products")
    routes.Post("/", h.AddProduct)
    routes.Get("/", h.GetProducts)
    routes.Get("/:cid", h.GetProduct)
    // routes.Put("/:id", h.UpdateBook)
    // routes.Delete("/:id", h.DeleteBook)
}