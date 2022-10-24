package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Project struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Location    Location           `json:"location" bson:"location" binding:"required"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}
type Location struct {
	StreetAddress string      `json:"streetaddress" bson:"streetaddress"`
	Lga           string      `json:"lga" bson:"lga" binding:"required"`
	State         string      `json:"state" bson:"state" binding:"required"`
	Country       string      `json:"country" bson:"country" binding:"required"`
	Coordinates   Coordinates `json:"coordinates" bson:"coordinates" binding:"required"`
}

type Coordinates struct {
	Longitude string `json:"longitude" bson:"longitude"`
	Latitude  string `json:"latitude" bson:"latitude" binding:"required"`
}

type Config struct{
	App_Port int
	Database_Host string
	Database_Port int
	}
