package mongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	logger "github.com/mixi-gaminh/core-framework/logs"
	mongoModel "github.com/mixi-gaminh/core-framework/model/mongodb"

	"gopkg.in/mgo.v2/bson"
)

//SortInMongo - SortInMongo
func (c *Mgo) SortInMongo(DBName string, collection, sortByField string) (interface{}, error) {
	var result []bson.M
	err := selectSession().DB(DBName).C(collection).Find(bson.M{}).Sort(sortByField).All(&result)
	if err != nil || len(result) == 0 {
		logger.ERROR(err)
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
			logger.ERROR(err)
			return nil, err
		}
	} else {
		err := selectSession().DB(DBName).C(collection).Find(bson.M{}).Sort(sortByField).All(&result)
		if err != nil || len(result) == 0 {
			logger.ERROR(err)
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
		logger.ERROR(err)
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
	findContent := bson.M{}
	if len(andQuery) > 0 {
		findContent = bson.M{"$and": andQuery}
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
	total, err := selectSession().DB(DBName).C(collection).Find(findContent).Count()
	if err != nil {
		logger.ERROR(err)
		return nil, 0, err
	}
	err = selectSession().DB(DBName).C(collection).Find(findContent).Sort(order + fieldSort).Skip(skip).Limit(limit).All(&retValue)
	if err != nil || len(retValue) == 0 {
		logger.ERROR(err)
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
		logger.ERROR(err)
		return nil, 0, err
	}
	err = selectSession().DB(DBName).C(collection).Find(bson.M{"$or": andQuery}).Sort(order + fieldSort).Skip(skip).Limit(limit).All(&retValue)
	if err != nil || len(retValue) == 0 {
		logger.ERROR(err)
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
		logger.ERROR(err)
		return nil, 0, err
	}
	err = selectSession().DB(DBName).C(collection).Find(bson.M{"$and": andQuery}).Sort(order + fieldSort).Skip(skip).Limit(limit).All(&retValue)
	if err != nil || len(retValue) == 0 {
		logger.ERROR(err)
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
		logger.ERROR(err)
		return nil, false
	}

	return retValue, true
}

// SearchInMongoByRange - Search Data By Range in MongoDB
func (c *Mgo) SearchInMongoByRange(DBName string, collection string, bodyBytes []byte) (interface{}, bool) {
	var bodyMap map[string]string
	json.Unmarshal(bodyBytes, &bodyMap)

	if bodyMap["field"] == "" {
		logger.ERROR("ERR: Field didn't Fill")
		return nil, false
	}
	field := bodyMap["field"]

	if bodyMap["gte"] == "" && bodyMap["gt"] == "" && bodyMap["lte"] == "" && bodyMap["lt"] == "" {
		logger.ERROR("ERR: GTE, LTE, GT, LT didn't Fill")
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
					logger.ERROR(err)
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
		logger.ERROR(err)
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
		logger.ERROR(err)
		return nil, 0, err
	}

	records := make([]interface{}, Limit)
	err = selectSession().DB(DBName).C(collection).Find(nil).Sort("_id").Skip(skip).Limit(limit).All(&records)
	if err != nil || len(records) == 0 {
		logger.ERROR(err)
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
		logger.ERROR(err)
		return nil, 0, err
	}

	if page != 0 && limit != 0 {
		err = selectSession().DB(DBName).C(collection).Find(query).Sort(sort...).Skip(skip).Limit(limit).All(&ret)
	} else {
		err = selectSession().DB(DBName).C(collection).Find(query).Sort(sort...).All(&ret)
	}

	if err != nil || len(ret) == 0 {
		logger.ERROR(err)
		return nil, 0, err
	}

	return ret, total, nil
}

// Statistics - Statistics
func (c *Mgo) Statistics(DBName string, collection string, inputData *mongoModel.Statistics) ([]map[string]interface{}, error) {
	match := map[string]interface{}{}
	group := map[string]interface{}{}

	// Initial conditions
	conditions := inputData.Conditions
	for _, item := range conditions {
		conditionItem := createConditionItem(item)
		match[item.FieldName] = conditionItem
	}

	// Initial groupby
	group["_id"] = createGroupBy(inputData.Result.GroupBy)
	fields := inputData.Result.Fields
	for _, item := range fields {
		group[item.FieldName] = map[string]interface{}{
			item.Operator: item.Column,
		}
	}

	others := inputData.Others
	sortField := others.Sort.Field
	sortType := others.Sort.Type
	limit := others.Limit

	// Execute query
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Pipe([]bson.M{{"$match": match}, {"$group": group}}).All(&result)
	if err != nil {
		return nil, err
	}
	if len(result) > 0 {
		for index, item := range result {
			// remove '_id' but keep it's properties
			keyGroup, OK := item["_id"].(map[string]interface{})
			if !OK {
				return nil, errors.New("_id parse failed")
			}
			for key, value := range keyGroup {
				result[index][key] = value
			}
			delete(result[index], "_id")
		}
		result, err = sortStatisticsResult(result, sortField, sortType, limit)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// BasicStatistics - BasicStatistics
func (c *Mgo) BasicStatistics(DBName string, collection string, match, project, group, sort map[string]interface{}, limit int) ([]map[string]interface{}, error) {
	// Initial
	pipeline := []bson.M{}
	if match != nil {
		pipeline = append(pipeline, bson.M{"$match": match})
	}
	if project != nil {
		pipeline = append(pipeline, bson.M{"$project": project})
	}
	if group != nil {
		pipeline = append(pipeline, bson.M{"$group": group})
	}
	if sort != nil {
		pipeline = append(pipeline, bson.M{"$sort": sort})
	}
	if limit > 0 {
		pipeline = append(pipeline, bson.M{"$limit": limit})
	}

	// Handle
	var result []map[string]interface{}
	err := selectSession().DB(DBName).C(collection).Pipe(pipeline).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
