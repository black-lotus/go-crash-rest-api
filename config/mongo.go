package config

import (
	"os"

	"gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	var host string = os.Getenv("MONGO_HOST")
	var dbName string = os.Getenv("MONGO_DB_NAME")

	var session *mgo.Session
	var err error
	session, err = mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}
