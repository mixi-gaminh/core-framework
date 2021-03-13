package mongo

import (
	"encoding/json"
	"strings"

	logger "github.com/mixi-gaminh/core-framework/logs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SaveMongoMQ - SaveMongoMQ
func (c *Mgo) SaveMongoMQ(DBName string, collection string, ID string, data map[string]interface{}) error {
	delete(data, "_id")
	_, err := selectSession().DB(DBName).C(collection).Upsert(bson.M{"_id": ID}, data)
	if err != nil {
		logger.ERROR("ERROR: ", DBName, collection, err)
		return err
	}
	return nil
}

// UpdateMongoMQ - UpdateMongoMQ
func (c *Mgo) UpdateMongoMQ(DBName string, collection, key string, data []byte) error {
	value := make(map[string]interface{})
	err := json.Unmarshal(data, &value)
	if err != nil {
		logger.ERROR(err)
		return err
	}

	delete(value, "_id")

	filter := bson.M{"_id": bson.M{"$eq": key}}
	update := bson.M{"$set": value}

	err = selectSession().DB(DBName).C(collection).UpdateId(filter, update)
	if err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// DeleteInMongoMQ - DeleteInMongoMQ
func (c *Mgo) DeleteInMongoMQ(DBName string, collection, deleteByID string) error {
	// Move Record to Recycle Bin
	persistData := c.FindByIDMQ(DBName, collection, deleteByID)
	if persistData == nil {
		logger.ERROR("Get Data for Move to Recycle Bin FAILED")
		return nil
	}
	delete(persistData, "_id")
	if _, err := selectSession().DB(MgoDBNameRecycleBin).C(collection).Upsert(bson.M{"_id": deleteByID}, persistData); err != nil {
		logger.ERROR(err)
		return err
	}
	if err := selectSession().DB(DBName).C(collection).Remove(bson.M{"_id": deleteByID}); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// ForceDeleteGeneral - ForceDeleteGeneral
func (c *Mgo) ForceDeleteGeneral(DBName string, collection, deleteByID string) error {
	if err := selectSession().DB(DBName).C(collection).Remove(bson.M{"_id": deleteByID}); err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// DropCollectionInMongoMQ - DropCollectionInMongoMQ
func (c *Mgo) DropCollectionInMongoMQ(DBName string, collection string) error {
	err := selectSession().DB(DBName).C(collection).DropCollection()
	if err != nil {
		logger.ERROR(err)
		return err
	}
	return nil
}

// DropManyCollectionInMongoMQ - DropManyCollectionInMongoMQ
func (c *Mgo) DropManyCollectionInMongoMQ(DBName string, listCollection []string) error {
	for _, collection := range listCollection {
		collection = strings.ReplaceAll(collection, "$", "@")
		err := selectSession().DB(DBName).C(collection).DropCollection()
		if err != nil {
			logger.ERROR("Error while drop collection ", collection)
			return err
		}
	}
	return nil
}

// FindAllInMongoMQ - FindAllInMongoMQ
func (c *Mgo) FindAllInMongoMQ(DBName string, collection string) []interface{} {
	var result []interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{}).All(&result)
	if err != nil {
		logger.ERROR(err)
		return nil
	}
	return result
}

//FindAllRegexByIDMQ - FindAllRegexByIDMQ
func (c *Mgo) FindAllRegexByIDMQ(DBName string, collection, id string) []map[string]interface{} {
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"_id": bson.M{"$regex": id}}).Sort("timestamp").All(&result)
	if err != nil {
		logger.ERROR(err)
		return nil
	}
	return result
}

// GetCollectionNamesMQ - GetCollectionNamesMQ
func (c *Mgo) GetCollectionNamesMQ(DBName string) ([]string, error) {
	collectionNames, err := selectSession().DB(DBName).CollectionNames()
	if err != nil {
		logger.ERROR(err)
		return nil, err
	}
	return collectionNames, nil
}

//FindByIDMQ - FindByIDMQ
func (c *Mgo) FindByIDMQ(DBName string, collection string, id string) map[string]interface{} {
	result := make(map[string]interface{})
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"_id": id}).One(&result)

	if err == mgo.ErrNotFound || err != nil {
		logger.ERROR(err)
		return nil
	}
	return result
}
