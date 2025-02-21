package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	AssignedTo  string             `bson:"assigned_to" json:"assigned_to"`
	DueDate     string             `bson:"due_date" json:"due_date"`
}
