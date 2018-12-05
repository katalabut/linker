package db

import (
	"github.com/katalabut/linker/utils"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
)

type docUrl struct {
	Id       bson.ObjectId `bson:"_id"`
	ShortUrl string        `bson:"short"`
	LongUrl  string        `bson:"url"`
}

func (c *MgoConnection) AddUrls(longUrl string) (link *docUrl, err error) {

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

	link = &docUrl{}

	err = urlCollection.Find(bson.M{"url": parseUrl.String()}).One(&link)
	if err != nil {
		link = &docUrl{
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

func (c *MgoConnection) FindlongUrl(shortUrl string) (lUrl string, err error) {
	//create an empty document struct
	result := docUrl{}
	//get a copy of the original session and a collection
	session, urlCollection, err := c.getSessionAndCollection()
	if err != nil {
		return
	}
	defer session.Close()
	//Find the shorturl that we need
	err = urlCollection.Find(bson.M{"short": shortUrl}).One(&result)
	if err != nil {
		return
	}
	return result.LongUrl, nil
}
