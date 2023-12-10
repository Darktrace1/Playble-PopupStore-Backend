package main

import (
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/cloud"
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/config"
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/db"
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/utils"
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/products"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()
	utils.CheckErr(err)

	app := fiber.New()
	db := db.Init(c.DBUrl)
	uploader := cloud.InitS3(c.AwsRegion, c.AwsAccessKey, c.AwsSecretKey)

	products.RegisterRoutes(app, db, uploader)

	app.Listen(c.Port)
}