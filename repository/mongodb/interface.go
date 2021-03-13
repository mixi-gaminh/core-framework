package mongo

import (
	model "github.com/mixi-gaminh/core-framework/model/mongodb"
	"gopkg.in/mgo.v2/bson"
)

// IMongoDB - IMongoDB
type IMongoDB interface {
	MongoDBConstructor()
	Close()
	Ping() error

	// SortInMongo - SortInMongo
	SortInMongo(string, string) (interface{}, error)

	// LimitSortInMongo - LimitSortInMongo
	LimitSortInMongo(string, string, int) ([]map[string]interface{}, error)

	// SearchManyConditionInMongo - Search Data with Many Condition in MongoDB
	SearchManyConditionInMongo(string, string, string, []string, []string) ([]interface{}, error)

	// SearchManyConditionInMongoWithPagging - Search Data with Many Condition with pagging in MongoDB
	SearchManyConditionInMongoWithPagging(string, string, string, []string, []string, int, int) (interface{}, int, error)

	// SearchManyConditionInMongoWithOrCondition - Search Data with Many Condition with OR condition
	SearchManyConditionInMongoWithOrCondition(string, string, string, []string, []string, int, int) (interface{}, int, error)

	// SearchNumericManyConditionWithPagging - Search Data Numeric with Many Condition with pagging in MongoDB
	SearchNumericManyConditionWithPagging(string, string, string, []string, []string, []float64, int, int) (interface{}, int, error)

	// SearchInMongo - Search Data in MongoDB
	SearchInMongo(string, string, string) (interface{}, bool)

	// SearchInMongoByRange - Search Data By Range in MongoDB
	SearchInMongoByRange(string, []byte) (interface{}, bool)

	//PaginateWithSkip - Phân trang dữ liệu.
	// VD: Page 1, Limit 2 => Page 1: record 1, record 2; Page 2: record 3, record 4
	//	   Page 2, Limit 3 => Page 1: record 1, record 2, record 3; Page 2: record 4, record 5, record 6
	PaginateWithSkip(string, int, int) ([]interface{}, int, error)

	// MatchDataByLogicClauseAndSort - MatchDataByLogicClauseAndSort
	MatchDataByLogicClauseAndSort(string, bson.M, []string, int, int) ([]interface{}, int, error)

	// LookUpInnerJoin - Use $lookup & $unwind & $project for Inner Join MongoDB
	LookUpInnerJoin(string, string, *model.InnerJoin, int, int) (interface{}, error)

	// GetTotalInnerJoin - Use $lookup & $unwind & $project for Inner Join MongoDB
	GetTotalInnerJoin(string, string, *model.InnerJoin) (interface{}, error)

	// CreateCollectionJSONSchema - CreateCollectionJSONSchema
	CreateCollectionJSONSchema(string, interface{}) error

	// CountDocuments - CountDocuments
	CountDocuments(string) (int, error)

	// GetCollectionNames - GetCollectionNames
	GetCollectionNames() ([]string, error)

	// FindByID - Find one document in Mongo DB by ID
	FindByID(string, string) map[string]interface{}

	// FindAll - Find all document in Mongo DB
	FindAll(string) []map[string]interface{}

	// FindAllID - Find all Document's ID in Mongo DB
	FindAllID(string) []map[string]interface{}

	// FindOneMongoByField - Find one document in Mongo DB by any Field
	FindOneMongoByField(string, string, string) (map[string]interface{}, bool)

	// SaveMongo - Save Data to Mongo DB
	SaveMongo(string, string, map[string]interface{}) error

	// UpdateManyByField - UpdateManyByField
	UpdateManyByField(string, string, []string, string, interface{}) error

	// UpdateMongo - UpdateMongo
	UpdateMongo([]byte, string, string)

	// DeleteInMongo - Delete Data in MongoDB
	DeleteInMongo(string, string) error

	// DeleteManyInMongo - Delete Many Data in MongoDB
	DeleteManyInMongo(string, ...string) error

	// DropCollectionInMongo - Delete a Collection in MongoDB
	DropCollectionInMongo(string)

	// DropManyCollectionInMongo - DropManyCollectionInMongo
	DropManyCollectionInMongo([]string)

	// DropDatabase - DropDatabase
	DropDatabase(string)

	// FindAllInMongo - FindAllInMongo
	FindAllInMongo(string) []string

	// FindAllRegexByID - Find All Data from Mongo DB by ID and Regex
	FindAllRegexByID(string, string) []map[string]interface{}

	// UpdateMany - UpdateMany
	UpdateMany(string, string, interface{}, ...string) error

	// SaveMongoMQ(collection string, ID string, data map[string]interface{})
	SaveMongoMQ(string, string, map[string]interface{}) error

	// UpdateMongoMQ(collection string, key string, data []byte) error
	UpdateMongoMQ(string, string, []byte) error

	// DeleteInMongoMQ(collection, deleteByID string) error
	DeleteInMongoMQ(string, string) error

	// ForceDeleteGeneral(dbName, collection, deleteByID string) error
	ForceDeleteGeneral(string, string, string) error

	// DropCollectionInMongoMQ(collection string) error
	DropCollectionInMongoMQ(string) error

	// DropManyCollectionInMongoMQ(listCollection []string) error
	DropManyCollectionInMongoMQ([]string) error

	// FindAllInMongoMQ(collection string) []interface{}
	FindAllInMongoMQ(string) []interface{}

	// FindAllRegexByIDMQ(collection, id string) []map[string]interface{}
	FindAllRegexByIDMQ(string, string) []map[string]interface{}

	// GetCollectionNamesMQ(DBName string) ([]string, error)
	GetCollectionNamesMQ(string) ([]string, error)

	// FindByIDMQ(collection string, id string) map[string]interface{}
	FindByIDMQ(string, string) map[string]interface{}

	// SaveMongoProtected(DBName string, collection string, ID string, data map[string]interface{}) error
	SaveMongoProtected(string, string, string, map[string]interface{}) error

	// FindProtectedByID(collection string, id string) map[string]interface{}
	FindProtectedByID(string, string) map[string]interface{}

	// FindProtectedAll(collection string) []map[string]interface{}
	FindProtectedAll(string) []map[string]interface{}
}
