package products

import (
	"context"
	"strings"
	"time"

	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/models"
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

type AddProductRequestBody struct {
	Name		string	`json:"name"`
	Provider	string	`json:"provider"`
	Price		int		`json:"price"`
	Image		string 	`json:"image,omitempty"`
}

func (h handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductRequestBody{}
	
	// 에러 처리
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	file, err := c.FormFile("file")
	utils.CheckErr(err)

	f, err := file.Open()
	utils.CheckErr(err)

	// 현재 시간을 기반으로 파일 이름 생성
	uniqueFileName := time.Now().Format("20060102150405") + "_" + file.Filename
	parts := strings.Split(file.Filename, ".")
	extension := parts[len(parts) - 1]

	// aws s3 upload
	// ContentDisposition -> inline
	// ContentType sets by file extension
	result, err := h.S3.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("jinbumnewtechbucket"),
		Key:    aws.String(uniqueFileName),
		Body:   f,
		ACL: 	"public-read",
		ContentDisposition: aws.String("inline"),
		ContentType: aws.String("image/" + extension),
	})
	utils.CheckErr(err)
	
	// body.Image에 S3 Link 추가
	body.Image = result.Location

	// Insert new db entry
	if result := h.DB.Create(context.TODO(), "new_tech", "products", &body); result != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error())
	}

	var product models.Product
		
	product.Id = h.DB.ID.String()
	product.Name = body.Name
	product.Price = body.Price
	product.Provider = body.Provider
	product.Image = result.Location

	return c.Status(fiber.StatusCreated).JSON(&product)
}