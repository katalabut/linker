package db

import "gopkg.in/mgo.v2/bson"

type Url struct {
	Id       bson.ObjectId `bson:"_id"`
	ShortUrl string        `bson:"short"`
	LongUrl  string        `bson:"url"`
}
