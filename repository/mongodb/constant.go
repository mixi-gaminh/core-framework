package mongo

import (
	"math/rand"
	"os"
	"time"

	logger "github.com/mixi-gaminh/core-framework/logs"
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
	logger.Constructor()
	logger.NewLogger()
	logger.INFO("MongoDB Constructor Successfull")
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
	sessionNumber := rand.Intn(maxSession-0+1) + 0
	return db[sessionNumber]
}
