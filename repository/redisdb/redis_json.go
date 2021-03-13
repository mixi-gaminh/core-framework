package redis

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"

	jonson "github.com/mixi-gaminh/core-framework/repository/redisdb/lib_jonson"
)

//Redis JSON v1: docker run -d -p 2222:6379 redislabs/rejson:latest
//Redis JSON v2: docker run -p 2222:6379 -d redislabs/redisjson2:0.8.0

var wg sync.WaitGroup

//ReJSONUpdate - JSON Redis Driver Update Data
func (c *Cache) ReJSONUpdate(ctx context.Context, key, path, value string) (string, error) {
	str, err := redisJSONWrite0.JsonSet(ctx, key, path, value).Result()
	if err != nil {
		log.Println(err)
		return str, err
	}
	return str, err
}

//ReJSONSet - JSON Redis Driver Set Data
func (c *Cache) ReJSONSet(ctx context.Context, key, path, value string, args interface{}) (string, error) {
	str, err := redisJSONWrite0.JsonSet(ctx, key, path, value, args).Result()
	if err != nil {
		log.Println(err)
		return str, err
	}
	return str, err
}

//ReJSONGetString - JSON Redis Driver Get a Data
func (c *Cache) ReJSONGetString(ctx context.Context, key string, query string) (string, error) {
	//Handing JSON Get with redis json v2
	if query == "" {
		query = "."
	}
	var args []interface{}
	args = append(args, "NOESCAPE")
	args = append(args, query)
	jsonString, err := redisJSONRead0.JsonGet(ctx, key, args...).Result()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return jsonString, nil
}

//ReJSONGet - JSON Redis Driver Get a Data
func (c *Cache) ReJSONGet(ctx context.Context, key string, query string) (interface{}, error) {
	if query == "" {
		query = "."
	}
	var args []interface{}
	args = append(args, "NOESCAPE")
	args = append(args, query)
	jsonString, err := redisJSONRead0.JsonGet(ctx, key, args...).Result()
	if err != nil {
		if !strings.Contains(err.Error(), "redis: nil") {
			log.Println(err)
		}
		return nil, err
	}

	var m interface{}
	err = json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return m, nil
}

func hashGetAllRoutine(waitgroup *sync.WaitGroup, jstring interface{}, key string, retArr []interface{}, i int, opts ...string) {
	defer waitgroup.Done()
	if jstring == nil {
		return
	}
	var getForType string = opts[0]
	var respData map[string]interface{}
	var m interface{}
	err := json.Unmarshal([]byte(jstring.(string)), &m)
	if err != nil {
		log.Println(err.Error())
		return
	}

	switch getForType {
	case "device":
		respData = map[string]interface{}{
			"device_id": strings.Split(key, "$")[2],
			"data_item": m,
		}
	case "bucket":
		if opts[1] == "" {
			return
		}
		respData = map[string]interface{}{
			"bucket_id": strings.Split(key, "$")[3],
			"device_id": strings.Split(key, "$")[2],
			"data_item": m,
		}

	case "bucket_all":
		respData = map[string]interface{}{
			"bucket_id": strings.Split(key, "$")[3],
			"device_id": strings.Split(key, "$")[2],
			"data_item": m,
		}
	case "record":
		if opts[1] == "" && opts[2] == "" {
			return
		}
		respData = map[string]interface{}{
			"record_id": strings.Split(key, "$")[3],
			"bucket_id": strings.Split(key, "$")[2],
			"device_id": strings.Split(key, "$")[1],
			"data_item": m,
		}
	case "record_all":
		respData = map[string]interface{}{
			"record_id": strings.Split(key, "$")[3],
			"bucket_id": strings.Split(key, "$")[2],
			"device_id": strings.Split(key, "$")[1],
			"data_item": m,
		}
	case "minio_record_all":
		respData = map[string]interface{}{
			"record_id": strings.Split(key, "$")[4],
			"bucket_id": strings.Split(key, "$")[3],
			"device_id": strings.Split(key, "$")[2],
			"data_item": m,
		}
	case "minio_bucket_all":
		respData = map[string]interface{}{
			"bucket_id": strings.Split(key, "$")[4],
			"device_id": strings.Split(key, "$")[3],
			"data_item": m,
		}
	default:
		respData = map[string]interface{}{
			"data_item": m,
		}
	}
	retArr[i] = respData
}

//HashGetAll - JSON Redis Driver Get All Data
func (c *Cache) HashGetAll(ctx context.Context, redisKey, query string, keys []string, opts ...string) (interface{}, error) {
	var waitgroup sync.WaitGroup

	jsonString, err := redisClientRead0.HMGet(ctx, redisKey, keys...).Result()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	retArr := make([]interface{}, len(jsonString))

	for i, jstring := range jsonString {
		waitgroup.Add(1)
		go hashGetAllRoutine(&waitgroup, jstring, keys[i], retArr, i, opts...)
	}
	waitgroup.Wait()

	retArr = deleteEmptyInArrInterface(retArr)
	return retArr, nil
}

//ReJSONGetAll - JSON Redis Driver Get All Data in Bucket
func (c *Cache) ReJSONGetAll(ctx context.Context, bucket, query string, keys []string, opts ...string) (interface{}, error) {
	var waitgroup sync.WaitGroup
	//var strArr []interface{}
	var resp responseAll
	var args []interface{}
	retArr := make([]interface{}, len(keys))

	var respData interface{}
	var getForType string = opts[0]

	args = append(args, "NOESCAPE")
	args = append(args, query)

	for i, k := range keys {
		waitgroup.Add(1)
		go func(i int, k string) {
			defer waitgroup.Done()
			jstring, err := redisJSONRead0.JsonGet(ctx, k, args...).Result()
			if err != nil && !strings.Contains(err.Error(), "redis: nil") {
				log.Println(err.Error())
				return
			}
			if jstring == "" {
				return
			}
			var m interface{}
			j, err := jonson.Parse([]byte(jstring))
			if err != nil {
				log.Println(err.Error())
			}
			m = j.ToInterface()

			switch getForType {
			case "device":
				respData = map[string]interface{}{
					"device_id": strings.Split(keys[i], "$")[2],
					"data_item": m,
				}

			case "bucket":
				if opts[1] == "" {
					return
				}
				respData = map[string]interface{}{
					"bucket_id": strings.Split(keys[i], "$")[3],
					"device_id": strings.Split(keys[i], "$")[2],
					"data_item": m,
				}
			case "bucket_all":
				respData = map[string]interface{}{
					"bucket_id": strings.Split(keys[i], "$")[3],
					"device_id": strings.Split(keys[i], "$")[2],
					"data_item": m,
				}
			case "record":
				if opts[1] == "" && opts[2] == "" {
					return
				}
				respData = map[string]interface{}{
					"record_id": strings.Split(keys[i], "$")[3],
					"bucket_id": strings.Split(keys[i], "$")[2],
					"device_id": strings.Split(keys[i], "$")[1],
					"data_item": m,
				}
			case "record_all":
				respData = map[string]interface{}{
					"record_id": strings.Split(keys[i], "$")[3],
					"bucket_id": strings.Split(keys[i], "$")[2],
					"device_id": strings.Split(keys[i], "$")[1],
					"data_item": m,
				}
			case "user_all":
				respData = map[string]interface{}{
					"record_id": strings.Split(keys[i], "$")[4],
					"bucket_id": strings.Split(keys[i], "$")[3],
					"device_id": strings.Split(keys[i], "$")[2],
					"data_item": m,
				}
			case "minio_record_all":
				respData = map[string]interface{}{
					"record_id": strings.Split(keys[i], "$")[4],
					"bucket_id": strings.Split(keys[i], "$")[3],
					"device_id": strings.Split(keys[i], "$")[2],
					"data_item": m,
				}
			case "minio_bucket_all":
				respData = map[string]interface{}{
					"bucket_id": strings.Split(keys[i], "$")[4],
					"device_id": strings.Split(keys[i], "$")[3],
					"data_item": m,
				}
			case "minio_id_all":
				respData = strings.Split(keys[i], "$")[4]
			default:
				respData = map[string]interface{}{
					"data_item": m,
				}
			}
			retArr[i] = respData
		}(i, k)
	}
	waitgroup.Wait()
	resp.Data = retArr
	resp.Data = deleteEmptyInArrInterface(resp.Data)
	return resp, nil
}

// ReJSONGetAnyAll - ReJSONGetAnyAll
func (c *Cache) ReJSONGetAnyAll(ctx context.Context, keys []string, query string, opts ...string) (interface{}, error) {
	firstKey := keys[0]
	var strArr []interface{}
	var args []interface{}

	for _, key := range keys[1:] {
		args = append(args, key)
	}

	args = append(args, "NOESCAPE")
	args = append(args, query)
	jsonString, err := redisJSONRead0.JsonMGet(ctx, firstKey, args...).Result()
	if err != nil {
		log.Println(err.Error())
		//return nil, err
	}
	for _, jstring := range jsonString {
		if jstring == "" {
			continue
		}
		var m interface{}
		err = json.Unmarshal([]byte(jstring), &m)
		if err != nil {
			log.Println(err.Error())
		}
		strArr = append(strArr, m)
	}

	return strArr, nil
}

//ReJSONArrLen - JSON Redis Driver Get Array Length
func (c *Cache) ReJSONArrLen(ctx context.Context, key, query string) (interface{}, error) {
	len, err := redisJSONRead0.JsonArrLen(ctx, key, query).Result()
	return len, err
}

//Delete - Delete a Record
func (c *Cache) Delete(ctx context.Context, key ...string) error {
	err := redisClientWrite0.Del(ctx, key...).Err()
	return err
}

//DeleteAll - DeleteAll
func (c *Cache) DeleteAll(ctx context.Context, key string) {
	keys := redisClientRead0.Keys(ctx, key+"*").Val()
	for _, key := range keys {
		redisClientWrite0.Del(ctx, key)
	}
}

//RealtimeNotification - Push Notification to Redis Hub
func (c *Cache) RealtimeNotification(ctx context.Context, channel string, msg interface{}) error {
	err := redisClientWrite0.Publish(ctx, channel, msg).Err()
	defer redisClientWrite0.Close()
	return err
}

//Decode - Decode
func Decode(jsonString string) (interface{}, error) {
	rejson, err := jonson.Parse([]byte(jsonString))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	ret, err := rejson.ToJSONString()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var m interface{}
	err = json.Unmarshal([]byte(ret), &m)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return m, nil
}

//ReJSONGetByQuery - JSON Redis Driver Get Records By Query
func (c *Cache) ReJSONGetByQuery(ctx context.Context, key, query string) []interface{} {
	i := 0
	keys := redisClientRead0.Keys(ctx, key+"*").Val()
	var result []interface{}
	for _, k := range keys {
		wg.Add(1)
		go func(key string) {
			ret, err := redisJSONRead0.JsonGet(ctx, key, "NOESCAPE", query).Result()
			if err == nil {
				m, err := Decode(ret)
				if err != nil {
					log.Println(err)
				}
				result = append(result, m)
				i++
			}
			wg.Done()
		}(k)
	}
	wg.Wait()
	log.Println(i)

	return result
}

// DeleteEmptyInArrInterface - Delete All Empty Element Interface in Array
func deleteEmptyInArrInterface(s []interface{}) []interface{} {
	var r []interface{}
	for _, str := range s {
		if str != nil {
			r = append(r, str)
		}
	}
	return r
}
