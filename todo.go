package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Completed bool          `bson:"completed" json:"completed"`
	Due       time.Time     `bson:"due" json:"due"`
}
