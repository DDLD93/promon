package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ddld93/promon/contractor/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB_Connect struct {
	Session *mongo.Client
}

var (
	database   = "promon"
	collection = "contractor"
)

func ConnectDB(host string, port int) *DB_Connect {
	URI := fmt.Sprintf("mongodb://%s:%v", host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return &DB_Connect{Session: client}
}

func (s *DB_Connect) AddOne(contractor model.Project) (*mongo.InsertOneResult, error) {
	contractor.Id = primitive.NewObjectID()
	contractor.CreatedAt = time.Now()
	result, err := s.Session.Database(database).Collection(collection).InsertOne(context.TODO(), contractor)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *DB_Connect) GetAll() ([]model.Project, error) {
	var contractor []model.Project
	result, err := s.Session.Database(database).Collection(collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = result.All(context.TODO(), &contractor)
	if err != nil {
		return nil, err
	}
	return contractor, nil
}

func (s *DB_Connect) GetOne(id primitive.ObjectID) (model.Project, error) {
	var contractor model.Project
	err := s.Session.Database(database).Collection(collection).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&contractor)
	if err != nil {
		return contractor, err
	}
	return contractor, nil
}

func (s *DB_Connect) Update(id primitive.ObjectID, data model.Project) *mongo.SingleResult {
	result := s.Session.Database(database).Collection(collection).FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, data)
	return result
}
func (s *DB_Connect) Delete(id primitive.ObjectID) *mongo.SingleResult {
	result := s.Session.Database(database).Collection(collection).FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	return result
}
