package mongoDB

import (
	"fmt"
	domain "gas-platform-service/internal/core/domain/gas-platform-service"
	"gas-platform-service/internal/core/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

func (m *MongoRepository) CreateCategory(platform *domain.Category) (*domain.Category, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Creating category with reference: %s ...", platform))
	coll := m.Client.Database("gasplus").Collection("category-collection")
	result, err := coll.InsertOne(context.TODO(), platform)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return platform, err
}

func (m *MongoRepository) UpdateCategory(platformReference string, platform *domain.Category) (*domain.Category, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Updating  category with reference: %s ...", platformReference))
	coll := m.Client.Database("gasplus").Collection("category-collection")
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
	helper.LogEvent("INFO", fmt.Sprintf("Getting category with reference: %s ...", reference))
	coll := m.Client.Database("gasplus").Collection("category-collection")
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
	helper.LogEvent("INFO", fmt.Sprintf("Getting category with name: %s ...", name))
	coll := m.Client.Database("gasplus").Collection("category-collection")
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
	coll := m.Client.Database("gasplus").Collection("category-collection")

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

	coll := m.Client.Database("gasplus").Collection("category-collection")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"reference": reference})

	if err != nil {
		helper.LogEvent("INFO", fmt.Sprintf("Deleting platform category with reference: %s failed.", reference))
		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	helper.LogEvent("INFO", fmt.Sprintf("Category with reference: %s deleted successfully", reference))
	return reference, nil
}

func (m *MongoRepository) CreateSubCategory(platform *domain.SubCategory) (*domain.SubCategory, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Creating sub category with sub-category : %s ...", platform))
	coll := m.Client.Database("gasplus").Collection("sub-category-collection")
	result, err := coll.InsertOne(context.TODO(), platform)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return platform, err
}

func (m *MongoRepository) UpdateSubCategory(reference string, platform *domain.SubCategory) (*domain.SubCategory, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Updating sub category with reference: %s ...", reference))
	coll := m.Client.Database("gasplus").Collection("sub-category-collection")
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
	helper.LogEvent("INFO", fmt.Sprintf("Getting sub category with reference: %s ...", reference))
	coll := m.Client.Database("gasplus").Collection("sub-category-collection")
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
	helper.LogEvent("INFO", fmt.Sprintf("Getting sub-category with name: %s ...", name))
	coll := m.Client.Database("gasplus").Collection("sub-category-collection")
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
	coll := m.Client.Database("gasplus").Collection("sub-category-collection")

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

	coll := m.Client.Database("gasplus").Collection("sub-category-collection")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"reference": reference})

	if err != nil {
		helper.LogEvent("INFO", fmt.Sprintf("Deleting platform category with reference: %s failed.", reference))
		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	helper.LogEvent("INFO", fmt.Sprintf("SubCategory with reference: %s deleted successfully", reference))
	return reference, nil
}

func (m *MongoRepository) CreateState(state *domain.State) (*domain.State, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Creating state category with reference: %s ...", state))
	coll := m.Client.Database("gasplus").Collection("state-collection")
	result, err := coll.InsertOne(context.TODO(), state)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return state, err
}

func (m *MongoRepository) UpdateState(stateReference string, state *domain.State) (*domain.State, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Updating state category with reference: %s ...", stateReference))
	coll := m.Client.Database("gasplus").Collection("state-collection")
	state.Reference = stateReference
	var updateDoc = &domain.State{}

	err := coll.FindOneAndUpdate(context.TODO(), bson.M{"reference": stateReference}, bson.D{{"$set", state}}).Decode(updateDoc)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return updateDoc, nil
}

func (m *MongoRepository) GetStateByReference(reference string) ([]*domain.State, error) {
	helper.LogEvent("INFO", fmt.Sprintf("Getting state category with reference: %s ...", reference))
	coll := m.Client.Database("gasplus").Collection("state-collection")
	filter := bson.D{{"reference", reference}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []*domain.State
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, err
}

func (m *MongoRepository) GetStateList(page int64) ([]*domain.State, error, int64) {

	coll := m.Client.Database("gasplus").Collection("state-collection")

	var results []*domain.State
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
		var elem *domain.State
		_ = cursor.Decode(&elem)
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("%v", err)
		return nil, err, 0
	}
	return results, err, totalPageCount
}

func (m *MongoRepository) DeleteStateByReference(reference string) (interface{}, error) {

	helper.LogEvent("INFO", fmt.Sprintf("Deleting sub category with reference: %s ...", reference))

	coll := m.Client.Database("gasplus").Collection("state-collection")
	_, err := coll.DeleteOne(context.TODO(), bson.M{"reference": reference})

	if err != nil {
		helper.LogEvent("INFO", fmt.Sprintf("Deleting platform category with reference: %s failed.", reference))
		return nil, helper.PrintErrorMessage("500", err.Error())
	}
	helper.LogEvent("INFO", fmt.Sprintf("SubCategory with reference: %s deleted successfully", reference))
	return reference, nil
}
