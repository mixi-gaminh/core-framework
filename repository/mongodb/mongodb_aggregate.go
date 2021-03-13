package mongo

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	model "github.com/mixi-gaminh/core-framework/model/mongodb"
	"gopkg.in/mgo.v2/bson"
)

//LookUpInnerJoin - Use $lookup & $unwind & $project for Inner Join MongoDB
func (c *Mgo) LookUpInnerJoin(DBName string, userID, deviceID string, innerJoin *model.InnerJoin, page, limit int) (interface{}, error) {
	prefixBucketID := "BM" + "@" + userID + "@" + deviceID + "@"
	var pipelines []bson.M
	for _, joinWith := range innerJoin.JoinWith {
		stageLookUp := bson.M{"$lookup": bson.M{
			"from":         prefixBucketID + joinWith.JoinBucketID,
			"localField":   joinWith.LocalField,
			"foreignField": joinWith.ForeignField,
			"as":           joinWith.JoinBucketID,
		}}
		stageUnWind := bson.M{"$unwind": "$" + joinWith.JoinBucketID}
		stageProject := bson.M{"$project": bson.M{
			"_id":                             0,
			"bucket":                          0,
			"createdByDevice":                 0,
			"created_at":                      0,
			"creatingDate":                    0,
			"description":                     0,
			joinWith.JoinBucketID + "._id":    0,
			joinWith.JoinBucketID + ".bucket": 0,
			joinWith.JoinBucketID + ".createdByDevice": 0,
			joinWith.JoinBucketID + ".created_at":      0,
			joinWith.JoinBucketID + ".creatingDate":    0,
			joinWith.JoinBucketID + ".description":     0,
			joinWith.JoinBucketID + ".id_incr":         0,
		}}
		if joinWith.Select != nil {
			temp := make(bson.M)
			for _, keySelect := range joinWith.Select {
				temp[keySelect] = 1
			}
			stageProject = bson.M{"$project": temp}
		}

		pipelines = append(pipelines, stageLookUp, stageUnWind)

		filter, err := handlerFilter(joinWith.Filter)
		if err != nil {
			return nil, err
		}
		if filter != nil {
			pipelines = append(pipelines, filter)
		}

		sort, err := handlerSort(joinWith.Sort)
		if err != nil {
			return nil, err
		}
		if sort != nil {
			pipelines = append(pipelines, sort)
		}
		pipelines = append(pipelines, stageProject)
	}

	// Operator In Handle
	for _, item := range innerJoin.OtherAggregate {
		pipelines = append(pipelines, bson.M{item.Operator: item.Defination})
	}

	// pageAndLimit := []bson.M{}
	if page != 0 && limit != 0 {
		skip := (page - 1) * limit
		skipQuery := bson.M{"$skip": skip}
		limitQuery := bson.M{"$limit": limit}
		pipelines = append(pipelines, skipQuery, limitQuery)
		// pageAndLimit = []bson.M{skipQuery, limitQuery}
	}

	// Create GroupBy
	group := map[string]interface{}{}
	initGroupBy := make(map[string]interface{})
	for _, item := range innerJoin.GroupBy.GroupID {
		initGroupBy[item.Key] = item.Value
	}
	if len(initGroupBy) > 0 {
		group["_id"] = initGroupBy
		statFields := innerJoin.GroupBy.StatField
		for _, item := range statFields {
			group[item.Field] = map[string]interface{}{
				item.Operator: item.Value,
			}
		}
		pipelines = append(pipelines, bson.M{"$group": group})
	}

	pipe := selectSession().DB(DBName).C(prefixBucketID + innerJoin.LocalBucketID).Pipe(pipelines).AllowDiskUse()
	var resp []interface{}
	err := pipe.All(&resp)
	if err != nil {
		log.Println("Errored: ", err)
		return nil, err
	}
	return resp, nil
}

// GetTotalInnerJoin - Use $lookup & $unwind & $project for Inner Join MongoDB
func (c *Mgo) GetTotalInnerJoin(DBName string, userID, deviceID string, innerJoin *model.InnerJoin) (interface{}, error) {
	prefixBucketID := "BM" + "@" + userID + "@" + deviceID + "@"
	var pipelines []bson.M
	for _, joinWith := range innerJoin.JoinWith {
		stageLookUp := bson.M{"$lookup": bson.M{
			"from":         prefixBucketID + joinWith.JoinBucketID,
			"localField":   joinWith.LocalField,
			"foreignField": joinWith.ForeignField,
			"as":           joinWith.JoinBucketID,
		}}
		stageUnWind := bson.M{"$unwind": "$" + joinWith.JoinBucketID}
		stageProject := bson.M{"$project": bson.M{
			"_id":                             0,
			"bucket":                          0,
			"createdByDevice":                 0,
			"created_at":                      0,
			"creatingDate":                    0,
			"description":                     0,
			joinWith.JoinBucketID + "._id":    0,
			joinWith.JoinBucketID + ".bucket": 0,
			joinWith.JoinBucketID + ".createdByDevice": 0,
			joinWith.JoinBucketID + ".created_at":      0,
			joinWith.JoinBucketID + ".creatingDate":    0,
			joinWith.JoinBucketID + ".description":     0,
			joinWith.JoinBucketID + ".id_incr":         0,
		}}
		if joinWith.Select != nil {
			temp := make(bson.M)
			for _, keySelect := range joinWith.Select {
				temp[keySelect] = 1
			}
			stageProject = bson.M{"$project": temp}
		}

		pipelines = append(pipelines, stageLookUp, stageUnWind)

		filter, err := handlerFilter(joinWith.Filter)
		if err != nil {
			return nil, err
		}
		if filter != nil {
			pipelines = append(pipelines, filter)
		}
		pipelines = append(pipelines, stageProject)
	}

	// Operator In Handle
	for _, item := range innerJoin.OtherAggregate {
		pipelines = append(pipelines, bson.M{item.Operator: item.Defination})
	}

	pipe := selectSession().DB(DBName).C(prefixBucketID + innerJoin.LocalBucketID).Pipe(pipelines).AllowDiskUse()
	var resp []interface{}
	err := pipe.All(&resp)
	if err != nil {
		log.Println("Errored: ", err)
		return nil, err
	}

	return len(resp), nil
}

func handlerFilter(data map[string]interface{}) (bson.M, error) {
	fields := make([]string, 0, len(data))
	values := make([]string, 0, len(data))
	for f, v := range data {
		if v == nil {
			log.Println("Filtering Value can not be nil")
			return nil, fmt.Errorf("FILTERING_VALUE_NIL")
		}
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			fields = append(fields, f)
			tmp := fmt.Sprintf("%s", v)
			values = append(values, tmp)
		case reflect.Slice:
			sliceValue, ok := v.([]interface{})
			if !ok {
				return nil, fmt.Errorf("Failed")
			}
			for _, v2 := range sliceValue {
				fields = append(fields, f)
				tmp := fmt.Sprintf("%s", v2)
				values = append(values, tmp)
			}
		default:
			return nil, fmt.Errorf("Failed")
		}
	}

	var andQuery []bson.M
	for i := 0; i < len(fields); i++ {
		//selector := bson.M{fields[i]: bson.M{"$regex": values[i]}}
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

	query := bson.M{"$match": bson.M{"$and": andQuery}}

	if len(andQuery) == 0 {
		return nil, nil
	}

	return query, nil
}

func handlerSort(listSort []model.SortFilterStruct) (bson.M, error) {
	if len(listSort) == 0 {
		return nil, nil
	}

	sort := make(bson.M)
	for _, val := range listSort {
		if strings.ToLower(val.Order) == "asc" {
			sort[val.Field] = 1
		} else if strings.ToLower(val.Order) == "des" {
			sort[val.Field] = -1
		} else {
			return nil, fmt.Errorf("SORT ORDER INVALID")
		}
	}

	return bson.M{"$sort": sort}, nil
}
