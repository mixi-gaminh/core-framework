package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SaveMongoProtected - Save Data to Mongo DB Protec
func (c *Mgo) SaveMongoProtected(DBName string, collection string, ID string, data map[string]interface{}) error {
	_, err := selectSession().DB(MgoDBNameProtected).C(collection).Upsert(bson.M{"_id": ID}, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//FindProtectedByID - Find one document in Mongo DB Protect by ID
func (c *Mgo) FindProtectedByID(DBName string, collection string, id string) map[string]interface{} {
	result := make(map[string]interface{})
	err := selectSession().DB(MgoDBNameProtected).C(collection).Find(bson.M{"_id": id}).One(&result)

	if err != mgo.ErrNotFound && err != nil {
		log.Println(err)
		return nil
	}
	return result
}

//FindProtectedAll - Find all document in Mongo DB Protect
func (c *Mgo) FindProtectedAll(DBName string, collection string) []map[string]interface{} {
	var result []map[string]interface{}
	err := selectSession().DB(MgoDBNameProtected).C(collection).Find(bson.M{}).All(&result)
	if err != mgo.ErrNotFound && err != nil {
		log.Println(err)
		return nil
	}
	return result
}
