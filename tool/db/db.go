package db

import (
	"errors"
	"gopkg.in/mgo.v2"
)

type MgoConnection struct {
	session *mgo.Session
}

func (c *MgoConnection) getSessionAndCollection() (session *mgo.Session, urlCollection *mgo.Collection, err error) {
	if c.session != nil {
		session = c.session.Copy()
		urlCollection = session.DB("db").C("collection")
	} else {
		err = errors.New("No original session found")
	}
	return
}

func (c *MgoConnection) AddUrls(longUrl string, shortUrl string) (err error) {
	//get a copy of the session
	/*session, urlCollection, err := c.getSessionAndCollection()
	if err == nil {
		defer session.Close()
		//insert a document with the provided function arguments
		err = urlCollection.Insert(
			&mongoDocument{
				Id:       bson.NewObjectId(),
				ShortUrl: shortUrl,
				LongUrl:  longUrl,
			},
		)
		if err != nil {
			//check if the error is due to duplicate shorturl
			if mgo.IsDup(err) {
				err = errors.New("Duplicate name exists for the shorturl")
			}
		}
	}*/
	return
}
