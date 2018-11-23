package db

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"linker/utils"
	"net/url"
)

type Url struct {
	Id       bson.ObjectId `bson:"_id"`
	ShortUrl string        `bson:"short"`
	LongUrl  string        `bson:"url"`
}

func (c *MgoConnection) AddUrls(longUrl string) (link *Url, err error) {

	parseUrl, err := url.ParseRequestURI(longUrl)
	if err != nil {
		err = errors.New("Link does not match format")
		return
	}

	session, urlCollection, err := c.getSessionAndCollection()
	if err != nil {
		err = errors.New("Session not found")
		return
	}

	defer session.Close()

	link = &Url{}

	err = urlCollection.Find(bson.M{"url": parseUrl.String()}).One(&link)
	if err != nil {
		link = &Url{
			Id:       bson.NewObjectId(),
			ShortUrl: utils.RandStringBytesMaskImprSrc(6),
			LongUrl:  parseUrl.String(),
		}

		err = urlCollection.Insert(link)
	}

	if err != nil {
		if mgo.IsDup(err) {
			err = errors.New("Duplicate name exists for the shorturl")
		}
	}
	return
}
