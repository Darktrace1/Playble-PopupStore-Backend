package db

import (
	"context"

	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/mongorm"
	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(url string) *mongorm.Model {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	utils.CheckErr(err)

	// 여기에 뭔가 MongoDB에 컬렉션을 &models.Product{} 에 대한 것으로 생성하는 코드가 들어가야할듯
	// db.auto
	m := &mongorm.Model{
		Client: db,
	}

	return m
}