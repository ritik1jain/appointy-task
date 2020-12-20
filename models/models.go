package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Dob       string             `json:"dob" bson:"dob,omitempty"`
	Phone     string             `json:"phone" bson:"phone,omitempty"`
	Timestamp time.Time         `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

// Create Contact Struct
type Contact struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserIDOne string             `json:"userid1,omitempty" bson:"userid1,omitempty"`
	UserIDTwo string             `json:"userid2,omitempty" bson:"userid2,omitempty"`
	Timestamp time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
