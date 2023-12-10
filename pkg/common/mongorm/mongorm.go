package mongorm

import (
	"context"
	"time"

	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	Client		*mongo.Client
	ID 			primitive.ObjectID 	`bson:"_id,omitempty"`
	CreatedAt 	time.Time 			`bson:"created_at"`
	UpdatedAt 	time.Time 			`bson:"updated_at"`
}

func (m *Model) Create(ctx context.Context, databaseName string, collectionName string, model interface{}) error {
	collection := m.Client.Database(databaseName).Collection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, model)
	utils.CheckErr(err)

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *Model) ReadAll(ctx context.Context, databaseName string, collectionName string, result interface{}) error {
	collection := m.Client.Database(databaseName).Collection(collectionName)
	
	// err := collection.Find(ctx, bson.M{}).Decode(result)
	// utils.CheckErr(err)
	
	// return nil
	// Find 메소드를 사용하여 커서를 얻습니다.
	cursor, err := collection.Find(ctx, bson.M{})
	utils.CheckErr(err)
	defer cursor.Close(ctx) // 함수 종료 전에 커서를 닫습니다.
	
	// All 메소드를 사용하여 모든 문서를 result에 디코딩합니다.
	if err := cursor.All(ctx, result); err != nil {
		return err // 에러 반환
	}
	
	return nil
}

func (m *Model) Read(ctx context.Context, databaseName string, collectionName string, filter interface{}, result interface{}) error {
	collection := m.Client.Database(databaseName).Collection(collectionName)
	
	err := collection.FindOne(ctx, filter).Decode(result)
	utils.CheckErr(err)
	
	return nil
}

func (m *Model) Update(ctx context.Context, databaseName string, collectionName string, filter interface{}, update interface{}) error {
	collection := m.Client.Database(databaseName).Collection(collectionName)

	m.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, filter, update)
	utils.CheckErr(err)

	return nil
}

func (m *Model) Delete(ctx context.Context, databaseName string, collectionName string, filter interface{}) error {
	collection := m.Client.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(ctx, filter)
	utils.CheckErr(err)

	return nil
}