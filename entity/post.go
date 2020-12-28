package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string        `bson:"title"`
	Text      string        `bson:"text"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
