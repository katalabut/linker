package db

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
)

type MgoConnection struct {
	session *mgo.Session
}

const Host = "192.168.99.100"
const Name = "linkerdb"

func NewDBConnection() (conn *MgoConnection) {
	conn = new(MgoConnection)
	conn.createLocalConnection()
	return
}

func (c *MgoConnection) createLocalConnection() (err error) {
	fmt.Println("Connecting to local mongo...")
	c.session, err = mgo.Dial(Host)

	if err != nil {
		fmt.Printf("Error occured while creating mongodb connection: %s", err.Error())
		return
	}

	fmt.Println("Connection established to mongo server")
	urlcollection := c.session.DB(Name).C("Urls")
	if urlcollection == nil {
		err = errors.New("Collection could not be created")
		return
	}

	index := mgo.Index{
		Key:      []string{"$text:short"},
		Unique:   true,
		DropDups: true,
	}
	urlcollection.EnsureIndex(index)
	return
}

func (c *MgoConnection) getSessionAndCollection() (session *mgo.Session, urlCollection *mgo.Collection, err error) {
	if c.session != nil {
		session = c.session.Copy()
		urlCollection = session.DB(Name).C("Urls")
	} else {
		err = errors.New("No original session found")
	}
	return
}
