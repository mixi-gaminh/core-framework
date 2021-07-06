package mongo

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"time"

	logger "github.com/mixi-gaminh/core-framework/logs"
	mongoModel "github.com/mixi-gaminh/core-framework/model/mongodb"
	"gopkg.in/mgo.v2"
)

// Mgo - Mongo DB Struct
type Mgo struct{}

// MgoDBNameProtected - MgoDBNameProtected
const MgoDBNameProtected = "vnpt_business_platform_protected"

// MgoDBNameRecycleBin - MgoDBNameRecycleBin
const MgoDBNameRecycleBin = "recyclebin_vnpt_business_platform"

//DB - DB
//var db *mgo.Session

//Limit - Limit
const Limit = 5

//TotalDocsNum - TotalDocsNum
const TotalDocsNum = 1000

// MongoHost - Mongo DB URL
var MongoHost []string

const maxSession int = 10

var db [maxSession]*mgo.Session

// MongoDBConstructor - MongoDBConstructor
func (c *Mgo) MongoDBConstructor(MongoHost []string, username, password string) {
	for i := 0; i < maxSession; i++ {
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    MongoHost,
			Timeout:  60 * time.Second,
			Username: username,
			Password: password,
		}
		_db, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			logger.ERROR("MONGO ESTABLISH CONNECTION ERROR: ", err)
			os.Exit(-1)
		}
		db[i] = _db
	}
	if err := c.Ping(); err != nil {
		logger.ERROR("MONGO PING ERROR: ", err)
		os.Exit(-1)
	}

	// Go routine check mongodb connection for reconnect machenism
	go c.checkSessionForReconnect(MongoHost, username, password)

	logger.NewLogger()
	logger.INFO("MongoDB Constructor Successfull")
}

func (c *Mgo) checkSessionForReconnect(MongoHost []string, username, password string) {
	for {
		// Check connection in every 60 seconds
		time.Sleep(60 * time.Second)

		// Initial
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    MongoHost,
			Timeout:  60 * time.Second,
			Username: username,
			Password: password,
		}

		// Handle
		for i := 0; i < maxSession; i++ {
			if err := db[i].Ping(); err != nil {
				fmt.Println("checkSessionForReconnect - Lost connection, session:", i)

				// Reconnect with new session
				_db, err := mgo.DialWithInfo(mongoDBDialInfo)
				if err != nil {
					fmt.Println("checkSessionForReconnect - Reconnect Failed, err:", err)
				} else {
					fmt.Println("checkSessionForReconnect - Reconnect successfully, session:", i)
					db[i] = _db
				}
			}
		}
	}
}

// Close - Close
func (c *Mgo) Close() {
	for i := 0; i < maxSession; i++ {
		db[i].Close()
	}
}

// Ping - Ping
func (c *Mgo) Ping() error {
	for i := 0; i < maxSession; i++ {
		if err := db[i].Ping(); err != nil {
			return err
		}
	}
	return nil
}

func isNumDot(s string) bool {
	dotFound := false
	for _, v := range s {
		if v == '.' {
			if dotFound {
				return false
			}
			dotFound = true
		} else if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func selectSession() *mgo.Session {
	sessionNumber := rand.Intn(maxSession)
	return db[sessionNumber]
}

func createConditionItem(item mongoModel.Conditions) map[string]interface{} {
	conditionItem := map[string]interface{}{}

	if item.Operators.Equal != nil {
		conditionItem["$eq"] = item.Operators.Equal
	}
	if item.Operators.GreaterThan != nil {
		conditionItem["$gt"] = item.Operators.GreaterThan
	}
	if item.Operators.LessThan != nil {
		conditionItem["$lt"] = item.Operators.LessThan
	}
	if item.Operators.GreaterThanOrEqual != nil {
		conditionItem["$gte"] = item.Operators.GreaterThanOrEqual
	}
	if item.Operators.LessThanOrEqual != nil {
		conditionItem["$lte"] = item.Operators.LessThanOrEqual
	}
	if item.Operators.NotEqual != nil {
		conditionItem["$ne"] = item.Operators.NotEqual
	}

	return conditionItem
}
func createGroupBy(groupBy []mongoModel.StatisticsGroupBy) map[string]string {
	result := make(map[string]string)
	for _, item := range groupBy {
		result[item.Key] = item.Value
	}
	return result
}

func sortStatisticsResult(data []map[string]interface{}, sortField, sortType string, limit int) ([]map[string]interface{}, error) {
	if data[0] == nil || data[0][sortField] == nil {
		return data, nil
	}

	if len(sortType) > 2 && sortType[0:3] == "ASC" {
		sort.SliceStable(data, func(i, j int) bool {
			ri := data[i][sortField]
			rj := data[j][sortField]
			if reflect.TypeOf(ri).String() == "string" {
				return ri.(string) < rj.(string)
			} else if reflect.TypeOf(ri).String() == "int64" {
				return ri.(int64) < rj.(int64)
			} else if reflect.TypeOf(ri).String() == "float64" {
				return ri.(float64) < rj.(float64)
			}
			return false
		})
	} else if len(sortType) > 2 && sortType[0:3] == "DES" {
		sort.SliceStable(data, func(i, j int) bool {
			ri := data[i][sortField]
			rj := data[j][sortField]
			if reflect.TypeOf(ri).String() == "string" {
				return ri.(string) > rj.(string)
			} else if reflect.TypeOf(ri).String() == "int64" {
				return ri.(int64) > rj.(int64)
			} else if reflect.TypeOf(ri).String() == "float64" {
				return ri.(float64) > rj.(float64)
			}
			return false
		})
	} else {
		return nil, errors.New("body invalid")
	}

	if limit < len(data) && limit > 0 {
		return data[:limit], nil
	}
	return data, nil
}
