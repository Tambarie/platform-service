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

func (m *MongoRepository) CreateCategory(platform *domain.Category) (*domain.Category, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	result, err := coll.InsertOne(context.TODO(), platform)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return platform, err
}

func (m *MongoRepository) UpdateCategory(platformReference string, platform *domain.Category) (*domain.Category, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	platform.Reference = platformReference
	var updateDoc = &domain.Category{}

	err := coll.FindOneAndUpdate(context.TODO(), bson.M{"reference": platformReference}, bson.D{{"$set", platform}}).Decode(updateDoc)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return updateDoc, nil
}

func (m *MongoRepository) GetCategoryByReference(reference string) ([]*domain.Category, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	filter := bson.D{{"reference", reference}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.Category
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetCategoryByName(name string) ([]*domain.Category, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	filter := bson.D{{"name", name}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.Category
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetCategoryList(page int64) ([]*domain.Category, error, int64) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")

	var results []*domain.Category
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
		var elem *domain.Category
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

func (m *MongoRepository) CreateSubCategory(platform *domain.SubCategory) (*domain.SubCategory, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	result, err := coll.InsertOne(context.TODO(), platform)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return platform, err
}

func (m *MongoRepository) UpdateSubCategory(reference string, platform *domain.SubCategory) (*domain.SubCategory, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	platform.Reference = reference
	var updateDoc = &domain.SubCategory{}

	err := coll.FindOneAndUpdate(context.TODO(), bson.M{"reference": reference}, bson.D{{"$set", platform}}).Decode(updateDoc)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return updateDoc, nil
}

func (m *MongoRepository) GetSubCategoryByReference(reference string) ([]*domain.SubCategory, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	filter := bson.D{{"reference", reference}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.SubCategory
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetSubCategoryByName(name string) ([]*domain.SubCategory, error) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	filter := bson.D{{"name", name}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.SubCategory
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetSubCategoryList(page int64) ([]*domain.SubCategory, error, int64) {
	coll := m.Client.Database("gasplus").Collection("gas-platform-service")

	var results []*domain.SubCategory
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
		var elem *domain.SubCategory
		_ = cursor.Decode(&elem)
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("%v", err)
		return nil, err, 0
	}
	return results, err, totalPageCount
}

func (m *MongoRepository) DeleteSubCategoryByReference(reference string) (interface{}, error) {

	helper.LogEvent("INFO", fmt.Sprintf("Deleting sub category with reference: %s ...", reference))

	coll := m.Client.Database("gasplus").Collection("gas-platform-service")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"reference": reference})

	if err != nil {
		helper.LogEvent("INFO", fmt.Sprintf("Deleting platform category with reference: %s failed.", reference))
		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	helper.LogEvent("INFO", fmt.Sprintf("SubCategory with reference: %s deleted successfully", reference))
	return reference, nil
}
