package model

import (
	"gopkg.in/mgo.v2/bson"
)

type CarPort struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"cartportid"`
	Name     string        `bson:"name" json:"name"`
	Location GeoJson       `bson:"location" json:"location"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

