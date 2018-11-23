package db

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Url struct {
	Id       bson.ObjectId `bson:"_id"`
	ShortUrl string        `bson:"short"`
	LongUrl  string        `bson:"url"`
}

func (c *MgoConnection) AddUrls(longUrl string) (url *Url, err error) {
	//get a copy of the session
	session, urlCollection, err := c.getSessionAndCollection()

	if err != nil {
		err = errors.New("Duplicate name exists for the shorturl")
		return
	}

	defer session.Close()

	url = &Url{
		Id:       bson.NewObjectId(),
		ShortUrl: "qwerty",
		LongUrl:  longUrl,
	}

	err = urlCollection.Insert(url)
	if err != nil {
		if mgo.IsDup(err) {
			err = errors.New("Duplicate name exists for the shorturl")
		}
	}
	return
}
