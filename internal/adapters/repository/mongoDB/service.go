package mongoDB

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	domain "platform-service/internal/core/domain/gas-platform-service"
	"platform-service/internal/core/helper"
)

func (m *MongoRepository) CreatePlatform(platform *domain.PlatformService) (*domain.PlatformService, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	result, err := coll.InsertOne(context.TODO(), platform)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return platform, err
}

func (m *MongoRepository) UpdatePlatform(platformReference string, platform *domain.PlatformService) (*domain.PlatformService, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	platform.Reference = platformReference
	var updateDoc = &domain.PlatformService{}

	err := coll.FindOneAndUpdate(context.TODO(), bson.M{"reference": platformReference}, bson.D{{"$set", platform}}).Decode(updateDoc)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return updateDoc, nil
}

func (m *MongoRepository) GetCategoryByReference(reference string) ([]*domain.PlatformService, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	filter := bson.D{{"reference", reference}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.PlatformService
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetCategoryByName(name string) ([]*domain.PlatformService, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	filter := bson.D{{"name", name}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.PlatformService
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetPlatformPage(page int64) ([]*domain.PlatformService, error, int64) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")

	var results []*domain.PlatformService
	var limit int64 = 10

	if page == 0 {
		page = 0
	}

	offset := (page - int64(1)) * limit
	opts := options.Find().SetSkip(offset).SetLimit(limit)
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	totalPageCount, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Printf("%v", err)
		return nil, err, 0
	}
	if err != nil {
		log.Printf("%v", err)
		return nil, err, 0
	}

	for cursor.Next(context.TODO()) {
		var elem *domain.PlatformService
		_ = cursor.Decode(&elem)
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("%v", err)
		return nil, err, 0
	}
	return results, err, totalPageCount
}

func (m *MongoRepository) DeleteCategoryByReference(reference string) (interface{}, error) {

	helper.LogEvent("INFO", fmt.Sprintf("Deleting platform category with reference: %s ...", reference))

	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"reference": reference})

	if err != nil {
		helper.LogEvent("INFO", fmt.Sprintf("Deleting platform category with reference: %s failed.", reference))
		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	helper.LogEvent("INFO", fmt.Sprintf("Category with reference: %s deleted successfully", reference))
	return reference, nil
}
