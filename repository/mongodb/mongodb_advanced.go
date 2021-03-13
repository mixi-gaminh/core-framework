package mongo

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//SortInMongo - SortInMongo
func (c *Mgo) SortInMongo(DBName string, collection, sortByField string) (interface{}, error) {
	var result []bson.M
	err := selectSession().DB(DBName).C(collection).Find(bson.M{}).Sort(sortByField).All(&result)
	if err != nil || len(result) == 0 {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

//LimitSortInMongo - LimitSortInMongo
func (c *Mgo) LimitSortInMongo(DBName string, collection, sortByField string, limit int) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	if sortByField == "" {
		err := selectSession().DB(DBName).C(collection).Find(bson.M{}).Limit(limit).All(&result)
		if err != nil || len(result) == 0 {
			log.Println(err)
			return nil, err
		}
	} else {
		err := selectSession().DB(DBName).C(collection).Find(bson.M{}).Sort(sortByField).All(&result)
		if err != nil || len(result) == 0 {
			log.Println(err)
			return nil, err
		}
	}
	return result, nil
}

//SearchManyConditionInMongo - Search Data with Many Condition in MongoDB
func (c *Mgo) SearchManyConditionInMongo(DBName string, bucket, fieldSort, orderSort string, fields, values []string) ([]interface{}, error) {
	var retValue []interface{}
	var andQuery []bson.M
	for i := 0; i < len(fields); i++ {
		//selector := bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		var selector bson.M
		if strings.HasPrefix(values[i], "$gte.") {
			values[i] = values[i][5:]
			selector = bson.M{fields[i]: bson.M{"$gte": values[i]}}
		} else if strings.HasPrefix(values[i], "$gt.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$gt": values[i]}}
		} else if strings.HasPrefix(values[i], "$lte.") {
			values[i] = values[i][5:]
			selector = bson.M{fields[i]: bson.M{"$lte": values[i]}}
		} else if strings.HasPrefix(values[i], "$lt.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$lt": values[i]}}
		} else if strings.HasPrefix(values[i], "$eq.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$eq": values[i]}}
		} else if strings.HasPrefix(values[i], "$ne.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$ne": values[i]}}
		} else {
			selector = bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		}
		andQuery = append(andQuery, selector)
	}

	if fieldSort == "" {
		fieldSort = "updated_at"
	}

	order := ""
	if strings.ToUpper(orderSort) == "ASC" {
		order = ""
	} else if strings.ToUpper(orderSort) == "DES" {
		order = "-"
	}

	err := selectSession().DB(DBName).C(bucket).Find(bson.M{"$and": andQuery}).Sort(order + fieldSort).All(&retValue)
	if err != nil || len(retValue) == 0 {
		log.Println(err)
		return nil, err
	}

	return retValue, nil
}

//SearchManyConditionInMongoWithPagging - Search Data with Many Condition with pagging in MongoDB
func (c *Mgo) SearchManyConditionInMongoWithPagging(DBName string, collection, fieldSort, orderSort string, fields, values []string, page, limit int) (interface{}, int, error) {
	if limit <= 0 {
		return nil, 0, fmt.Errorf("LIMIT MUST > 0")
	}

	var retValue []interface{}
	var andQuery []bson.M
	for i := 0; i < len(fields); i++ {
		//selector := bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		var selector bson.M
		if strings.HasPrefix(values[i], "$gte.") {
			values[i] = values[i][5:]
			selector = bson.M{fields[i]: bson.M{"$gte": values[i]}}
		} else if strings.HasPrefix(values[i], "$gt.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$gt": values[i]}}
		} else if strings.HasPrefix(values[i], "$lte.") {
			values[i] = values[i][5:]
			selector = bson.M{fields[i]: bson.M{"$lte": values[i]}}
		} else if strings.HasPrefix(values[i], "$lt.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$lt": values[i]}}
		} else if strings.HasPrefix(values[i], "$eq.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$eq": values[i]}}
		} else if strings.HasPrefix(values[i], "$ne.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$ne": values[i]}}
		} else {
			selector = bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		}
		andQuery = append(andQuery, selector)
	}

	if fieldSort == "" {
		fieldSort = "updated_at"
	}

	order := ""
	if strings.ToUpper(orderSort) == "ASC" {
		order = ""
	} else if strings.ToUpper(orderSort) == "DES" {
		order = "-"
	}

	skip := (page - 1) * limit
	total, err := selectSession().DB(DBName).C(collection).Find(bson.M{"$and": andQuery}).Count()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	err = selectSession().DB(DBName).C(collection).Find(bson.M{"$and": andQuery}).Sort(order + fieldSort).Skip(skip).Limit(limit).All(&retValue)
	if err != nil || len(retValue) == 0 {
		log.Println(err)
		return nil, 0, err
	}

	return retValue, total, nil
}

//SearchManyConditionInMongoWithOrCondition - Search Data with Many Condition with OR condition
func (c *Mgo) SearchManyConditionInMongoWithOrCondition(DBName string, collection, fieldSort, orderSort string, fields, values []string, page, limit int) (interface{}, int, error) {
	if limit <= 0 {
		return nil, 0, fmt.Errorf("LIMIT MUST > 0")
	}

	var retValue []interface{}
	var andQuery []bson.M
	for i := 0; i < len(fields); i++ {
		//selector := bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		var selector bson.M
		if strings.HasPrefix(values[i], "$gte.") {
			values[i] = values[i][5:]
			selector = bson.M{fields[i]: bson.M{"$gte": values[i]}}
		} else if strings.HasPrefix(values[i], "$gt.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$gt": values[i]}}
		} else if strings.HasPrefix(values[i], "$lte.") {
			values[i] = values[i][5:]
			selector = bson.M{fields[i]: bson.M{"$lte": values[i]}}
		} else if strings.HasPrefix(values[i], "$lt.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$lt": values[i]}}
		} else if strings.HasPrefix(values[i], "$eq.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$eq": values[i]}}
		} else if strings.HasPrefix(values[i], "$ne.") {
			values[i] = values[i][4:]
			selector = bson.M{fields[i]: bson.M{"$ne": values[i]}}
		} else {
			selector = bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		}
		andQuery = append(andQuery, selector)
	}

	if fieldSort == "" {
		fieldSort = "updated_at"
	}

	order := ""
	if strings.ToUpper(orderSort) == "ASC" {
		order = ""
	} else if strings.ToUpper(orderSort) == "DES" {
		order = "-"
	}

	skip := (page - 1) * limit
	total, err := selectSession().DB(DBName).C(collection).Find(bson.M{"$or": andQuery}).Count()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	err = selectSession().DB(DBName).C(collection).Find(bson.M{"$or": andQuery}).Sort(order + fieldSort).Skip(skip).Limit(limit).All(&retValue)
	if err != nil || len(retValue) == 0 {
		log.Println(err)
		return nil, 0, err
	}

	return retValue, total, nil
}

//SearchNumericManyConditionWithPagging - Search Data Numeric with Many Condition with pagging in MongoDB
func (c *Mgo) SearchNumericManyConditionWithPagging(DBName string, collection, fieldSort, orderSort string, fields, operators []string, values []float64, page, limit int) (interface{}, int, error) {
	if limit <= 0 {
		return nil, 0, fmt.Errorf("LIMIT MUST > 0")
	}
	var retValue []interface{}
	var andQuery []bson.M
	for i := 0; i < len(fields); i++ {
		//selector := bson.M{fields[i]: bson.M{"$regex": values[i], "$options": "i"}}
		var selector bson.M
		if strings.HasPrefix(operators[i], "$gte") {
			selector = bson.M{fields[i]: bson.M{"$gte": values[i]}}
		} else if strings.HasPrefix(operators[i], "$gt") {
			selector = bson.M{fields[i]: bson.M{"$gt": values[i]}}
		} else if strings.HasPrefix(operators[i], "$lte") {
			selector = bson.M{fields[i]: bson.M{"$lte": values[i]}}
		} else if strings.HasPrefix(operators[i], "$lt") {
			selector = bson.M{fields[i]: bson.M{"$lt": values[i]}}
		} else if strings.HasPrefix(operators[i], "$eq") {
			selector = bson.M{fields[i]: bson.M{"$eq": values[i]}}
		} else if strings.HasPrefix(operators[i], "$ne") {
			selector = bson.M{fields[i]: bson.M{"$ne": values[i]}}
		} else {
			return nil, 0, fmt.Errorf("BODY IS INVALID")
		}
		andQuery = append(andQuery, selector)
	}

	if fieldSort == "" {
		fieldSort = "updated_at"
	}

	order := ""
	if strings.ToUpper(orderSort) == "ASC" {
		order = ""
	} else if strings.ToUpper(orderSort) == "DES" {
		order = "-"
	}

	skip := (page - 1) * limit
	total, err := selectSession().DB(DBName).C(collection).Find(bson.M{"$and": andQuery}).Count()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	err = selectSession().DB(DBName).C(collection).Find(bson.M{"$and": andQuery}).Sort(order + fieldSort).Skip(skip).Limit(limit).All(&retValue)
	if err != nil || len(retValue) == 0 {
		log.Println(err)
		return nil, 0, err
	}

	return retValue, total, nil
}

//SearchInMongo - Search Data in MongoDB
func (c *Mgo) SearchInMongo(DBName string, collection, field, value string) (interface{}, bool) {
	var retValue []bson.M
	selector := bson.M{field: bson.M{"$regex": value, "$options": "i"}}
	err := selectSession().DB(DBName).C(collection).Find(selector).All(&retValue)
	if err != nil || len(retValue) == 0 {
		log.Println(err)
		return nil, false
	}

	return retValue, true
}

// SearchInMongoByRange - Search Data By Range in MongoDB
func (c *Mgo) SearchInMongoByRange(DBName string, collection string, bodyBytes []byte) (interface{}, bool) {
	var bodyMap map[string]string
	json.Unmarshal(bodyBytes, &bodyMap)

	if bodyMap["field"] == "" {
		log.Println("ERR: Field didn't Fill")
		return nil, false
	}
	field := bodyMap["field"]

	if bodyMap["gte"] == "" && bodyMap["gt"] == "" && bodyMap["lte"] == "" && bodyMap["lt"] == "" {
		log.Println("ERR: GTE, LTE, GT, LT didn't Fill")
		return nil, false
	}

	if bodyMap["gte"] != "" && bodyMap["gt"] != "" {
		return nil, false
	}
	if bodyMap["lte"] != "" && bodyMap["lt"] != "" {
		return nil, false
	}

	var andQuery []bson.M
	for key, val := range bodyMap {
		if key == "gte" || key == "gt" || key == "lte" || key == "lt" {
			var selector bson.M
			if !isNumDot(val) {
				selector = bson.M{field: bson.M{"$" + key: val}}
			} else {
				valInt, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					log.Println(err)
					return nil, false
				}
				selector = bson.M{field: bson.M{"$" + key: valInt}}

			}
			andQuery = append(andQuery, selector)
		}
	}

	var retValue []bson.M
	err := selectSession().DB(DBName).C(collection).Find(bson.M{"$and": andQuery}).All(&retValue)
	if err != nil || len(retValue) == 0 {
		log.Println(err)
		return nil, false
	}

	return retValue, true
}

//PaginateWithSkip - Phân trang dữ liệu.
// VD: Page 1, Limit 2 => Page 1: record 1, record 2; Page 2: record 3, record 4
//	   Page 2, Limit 3 => Page 1: record 1, record 2, record 3; Page 2: record 4, record 5, record 6
func (c *Mgo) PaginateWithSkip(DBName string, collection string, page, limit int) ([]interface{}, int, error) {
	if limit <= 0 {
		return nil, 0, fmt.Errorf("LIMIT MUST > 0")
	}

	start := time.Now()
	skip := (page - 1) * limit

	total, err := selectSession().DB(DBName).C(collection).Find(nil).Count()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	records := make([]interface{}, Limit)
	err = selectSession().DB(DBName).C(collection).Find(nil).Sort("_id").Skip(skip).Limit(limit).All(&records)
	if err != nil || len(records) == 0 {
		log.Println(err)
		return nil, 0, err
	}
	rFirstMap := records[0].(bson.M)
	rLastMap := records[len(records)-1].(bson.M)
	rFirst := fmt.Sprintf("%s", rFirstMap["_id"])
	rLast := fmt.Sprintf("%s", rLastMap["_id"])

	log.Printf("paginateWithSkip -> page %d with record from %s to %s in %s\n", page, rFirst, rLast, time.Since(start))
	return records, total, nil
}

// MatchDataByLogicClauseAndSort - MatchDataByLogicClauseAndSort
func (c *Mgo) MatchDataByLogicClauseAndSort(DBName string, collection string, query bson.M, sort []string, page, limit int) ([]interface{}, int, error) {
	var ret []interface{}
	skip := (page - 1) * limit
	total, err := selectSession().DB(DBName).C(collection).Find(query).Count()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	if page != 0 && limit != 0 {
		err = selectSession().DB(DBName).C(collection).Find(query).Sort(sort...).Skip(skip).Limit(limit).All(&ret)
	} else {
		err = selectSession().DB(DBName).C(collection).Find(query).Sort(sort...).All(&ret)
	}

	if err != nil || len(ret) == 0 {
		log.Println(err)
		return nil, 0, err
	}

	return ret, total, nil
}
