package mongo

import (
	"encoding/json"
	"fmt"

	"log"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateCollectionJSONSchema - CreateCollectionJSONSchema
func (c *Mgo) CreateCollectionJSONSchema(DBName string, collection string, jsonSchema interface{}) error {
	info := new(mgo.CollectionInfo)
	info.Validator = bson.M{"$jsonSchema": jsonSchema}
	err := selectSession().DB(DBName).C(collection).Create(info)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CountDocuments - CountDocuments
func (c *Mgo) CountDocuments(DBName string, collection string) (int, error) {
	n, err := selectSession().DB(DBName).C(collection).Count()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return n, nil
}

// GetCollectionNames - GetCollectionNames
func (c *Mgo) GetCollectionNames(DBName string) ([]string, error) {
	n, err := selectSession().DB(DBName).CollectionNames()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return n, nil
}

//FindByID - Find one document in Mongo DB by ID
func (c *Mgo) FindByID(DBName string, collection string, id string) map[string]interface{} {
	result := make(map[string]interface{})
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"_id": id}).One(&result)

	if err == mgo.ErrNotFound || err != nil {
		log.Println(err)
		return nil
	}
	return result
}

//FindAll - Find all document in Mongo DB
func (c *Mgo) FindAll(DBName string, collection string) []map[string]interface{} {
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{}).All(&result)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

//FindAllID - Find all Document's ID in Mongo DB
func (c *Mgo) FindAllID(DBName string, collection string) []map[string]interface{} {
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{}).Select(bson.M{"_id": "1"}).All(&result)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

//FindOneMongoByField - Find one document in Mongo DB by any Field
func (c *Mgo) FindOneMongoByField(DBName string, collection, f, v string) (map[string]interface{}, bool) {
	result := make(map[string]interface{})
	selector := bson.M{f: bson.M{"$regex": v}}
	err := selectSession().DB(DBName).C(collection).Find(selector).One(result)

	if err == mgo.ErrNotFound {
		return nil, false
	}
	return result, true
}

//SaveMongo - Save Data to Mongo DB
func (c *Mgo) SaveMongo(DBName string, collection string, ID string, data interface{}) error {
	_, err := selectSession().DB(DBName).C(collection).Upsert(bson.M{"_id": ID}, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// UpdateManyByField - UpdateManyByField
func (c *Mgo) UpdateManyByField(DBName string, collection string, fieldNameFilter string, valueFilter []string, fieldNameUpdate string, valueUpdate interface{}) error {
	filter := bson.M{fieldNameFilter: bson.M{"$in": valueFilter}}
	update := bson.M{"$set": bson.M{fieldNameUpdate: valueUpdate}}

	_, err := selectSession().DB(DBName).C(collection).UpdateAll(filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//UpdateMongo - UpdateMongo
func (c *Mgo) UpdateMongo(DBName string, collection string, data []byte, bucket, key string) {
	value := make(map[string]interface{})
	err := json.Unmarshal(data, &value)
	if err != nil {
		log.Println(err)
	}

	filter := bson.M{"id": bson.M{"$eq": key}}
	update := bson.M{"$set": value}

	err = selectSession().DB(DBName).C(bucket).UpdateId(filter, update)
	if err != nil {
		log.Println(err)
	}
}

//DeleteInMongo - Delete Data in MongoDB
func (c *Mgo) DeleteInMongo(DBName string, collection string, deleteByID string) error {
	data := make(map[string]interface{})
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"_id": deleteByID}).One(&data)
	if err == mgo.ErrNotFound || err != nil {
		return err
	}
	_, err = selectSession().DB(DBName).C("recyclebin_vnpt_business_platform").Upsert(bson.M{"_id": deleteByID}, data)
	if err != nil {
		log.Println(err)
		return err
	}
	err = selectSession().DB(DBName).C(collection).Remove(bson.M{"_id": deleteByID})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//DeleteManyInMongo - Delete Many Data in MongoDB
func (c *Mgo) DeleteManyInMongo(DBName string, collection string, listDocument ...string) error {
	_, err := selectSession().DB(DBName).C(collection).RemoveAll(bson.M{"_id": bson.M{"$in": listDocument}})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//DropCollectionInMongo - Delete a Collection in MongoDB
func (c *Mgo) DropCollectionInMongo(DBName string, collection string) {
	err := selectSession().DB(DBName).C(collection).DropCollection()
	if err != nil {
		log.Println(err)
	}
}

//DropManyCollectionInMongo - DropManyCollectionInMongo
func (c *Mgo) DropManyCollectionInMongo(DBName string, listCollection []string) {
	for _, collection := range listCollection {
		collection = strings.ReplaceAll(collection, "$", "@")
		err := selectSession().DB(DBName).C(collection).DropCollection()
		if err != nil {
			log.Println("Error while drop collection ", collection)
		}
	}
}

// DropDatabase - DropDatabase
func (c *Mgo) DropDatabase(DBNameDrop string) {
	err := selectSession().DB(DBNameDrop).DropDatabase()
	if err != nil {
		log.Println("Error while drop database ", DBNameDrop)
	}
}

// FindAllInMongo - FindAllInMongo
func (c *Mgo) FindAllInMongo(DBName string, collection string) []string {
	var result []bson.M
	var retData []string
	err := selectSession().DB(DBName).C(collection).Find(bson.M{}).All(&result)
	if err != nil {
		log.Println(err)
	}
	for _, r := range result {
		rTmp := fmt.Sprintf("%s", r["_id"])
		retData = append(retData, rTmp)
	}
	return retData
}

//FindAllRegexByID - Find All Data from Mongo DB by ID and Regex
func (c *Mgo) FindAllRegexByID(DBName string, collection, id string) []map[string]interface{} {
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"_id": bson.M{"$regex": id}}).Sort("timestamp").All(&result)
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

// UpdateMany - UpdateMany
func (c *Mgo) UpdateMany(DBName string, collection, fieldUpdate string, valueUpdate interface{}, listID ...string) error {
	selector := bson.M{"_id": bson.M{"$in": listID}}
	update := bson.M{"$set": bson.M{fieldUpdate: valueUpdate}}
	_, err := selectSession().DB(DBName).C(collection).UpdateAll(selector, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
