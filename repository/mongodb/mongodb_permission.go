package mongo

import (
	"fmt"

	model "github.com/mixi-gaminh/core-framework/model/authentication"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CountDocumentByField - CountDocumentByField
func (c *Mgo) CountDocumentByField(DBName string, collection, field, value string) (int, error) {
	n, err := selectSession().DB(DBName).C(collection).Find(bson.M{field: value}).Count()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return n, nil
}

//UserIsExist - Check User is Exists
func (c *Mgo) UserIsExist(DBName string, collection string, field, value string) bool {
	count, err := selectSession().DB(DBName).C(collection).Find(bson.M{field: value}).Count()
	if err == mgo.ErrNotFound || err != nil || count < 1 {
		return false
	}
	return true
}

// GetUserProfile - GetUserProfile
func (c *Mgo) GetUserProfile(DBName string, collection string, userID string) (map[string]interface{}, bool) {
	userProfile := make(map[string]interface{})
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"user_id": userID}).One(&userProfile)
	if err == mgo.ErrNotFound || err != nil {
		return userProfile, false
	}
	delete(userProfile, "password")
	delete(userProfile, "_id")
	return userProfile, true
}

//SaveMongoUniqueID - Save Data to Mongo DB with default unique _id
func (c *Mgo) SaveMongoUniqueID(DBName string, collection string, data interface{}) error {
	err := selectSession().DB(DBName).C(collection).Insert(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// UpdateUserProfleByUserID - UpdateUserProfleByUserID
func (c *Mgo) UpdateUserProfleByUserID(DBName string, collection, userID string, dataUpdate map[string]interface{}) error {
	filter := bson.M{"user_id": bson.M{"$eq": userID}}
	update := bson.M{"$set": dataUpdate}

	err := selectSession().DB(DBName).C(collection).Update(filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// RetrieveManyByField - RetrieveManyByField
func (c *Mgo) RetrieveManyByField(DBName string, collection string, fieldNameFilter string, valueFilter []string) []interface{} {
	var result []interface{}
	filter := bson.M{fieldNameFilter: bson.M{"$in": valueFilter}}
	selectSession().DB(DBName).C(collection).Find(filter).All(&result)
	return result
}

// UpdateByField - UpdateByField
func (c *Mgo) UpdateByField(DBName string, collection string, fieldKey string, fieldValue interface{}, data map[string]interface{}) error {
	filter := bson.M{fieldKey: bson.M{"$eq": fieldValue}}
	update := make(bson.M)
	for key, val := range data {
		update[key] = val
	}
	update = bson.M{"$set": update}

	err := selectSession().DB(DBName).C(collection).Update(filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//DeleteInMongoByField - Delete Data by field in MongoDB
func (c *Mgo) DeleteInMongoByField(DBName string, collection, field string, recordID interface{}) error {
	err := selectSession().DB(DBName).C(collection).Remove(bson.M{field: recordID})
	if err != nil {
		return err
	}
	return nil
}

// FindByUsername - FindByUsername
func (c *Mgo) FindByUsername(DBName string, collection string, username string) *model.UserProfile {
	result := new(model.UserProfile)
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"username": username}).One(&result)

	if err == mgo.ErrNotFound || err != nil {
		return nil
	}
	return result
}

// FindOneByField - FindOneByField
func (c *Mgo) FindOneByField(DBName string, collection string, fieldKey string, fieldValue interface{}) map[string]interface{} {
	var result map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{fieldKey: fieldValue}).One(&result)

	if err == mgo.ErrNotFound || err != nil {
		return nil
	}
	return result
}

// FindUserByField - FindUserByField
func (c *Mgo) FindUserByField(DBName string, collection string, fieldKey string, fieldValue interface{}) map[string]interface{} {
	var result map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(bson.M{fieldKey: fieldValue}).One(&result)
	if err == mgo.ErrNotFound || err != nil {
		fmt.Println(err)
		return nil
	}
	return result
}

// UpsertOneByField - UpsertOneByField
func (c *Mgo) UpsertOneByField(DBName string, collection string, fieldKey string, fieldValue interface{}, data interface{}) error {
	_, err := selectSession().DB(DBName).C(collection).Upsert(bson.M{fieldKey: fieldValue}, data)
	if err != nil {
		return err
	}
	return nil
}

// GetAllRecordInCollection - GetAllRecordInCollection
func (c *Mgo) GetAllRecordInCollection(DBName string, collection string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(nil).All(&result)

	if err == mgo.ErrNotFound || err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllRecordWithPagination - GetAllRecordWithPagination
func (c *Mgo) GetAllRecordWithPagination(DBName string, collection string, page int, pageSize int) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Find(nil).Skip((page - 1) * pageSize).Limit(pageSize).All(&result)
	if err == mgo.ErrNotFound || err != nil {
		return nil, err
	}
	return result, nil
}

// CountOfRows - CountOfRows
func (c *Mgo) CountOfRows(DBName string, collection string) (int, error) {
	result, err := selectSession().DB(DBName).C(collection).Count()
	if err == mgo.ErrNotFound || err != nil {
		return 0, err
	}
	return result, nil
}
