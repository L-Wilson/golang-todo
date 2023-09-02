package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// creating a model to define how the data will get stored in the DB:
// use struct in golang
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}
