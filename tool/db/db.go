package db

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
)

type MgoConnection struct {
	session *mgo.Session
}

func NewDBConnection() (conn *MgoConnection) {
	conn = new(MgoConnection)
	conn.createLocalConnection()
	return
}

func (c *MgoConnection) createLocalConnection() (err error) {
	fmt.Println("Connecting to local mongo...")
	c.session, err = mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Printf("Error occured while creating mongodb connection: %s", err.Error())
		return
	}

	fmt.Println("Connection established to mongo server")
	urlcollection := c.session.DB("LinkShortnerDB").C("UrlCollection")
	if urlcollection == nil {
		//err = errors.New("Collection could not be created, maybe need to create it manually")
		return
	}
	//This will create a unique index to ensure that there won't be duplicate shorturls in the database.
	index := mgo.Index{
		Key:      []string{"$text:shorturl"},
		Unique:   true,
		DropDups: true,
	}
	urlcollection.EnsureIndex(index)

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

func (c *MgoConnection) AddUrls(longUrl string) (err error) {
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
