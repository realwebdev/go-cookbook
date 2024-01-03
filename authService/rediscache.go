package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"strconv"
// 	"sync"

// 	"github.com/azure/et-3.0/com.code.et/ginmodels/models"
// 	"github.com/azure/et-3.0/com.code.et/helper/config"
// 	"github.com/go-redsync/redsync/v4"
// 	"github.com/go-redsync/redsync/v4/redis/redigo"
// 	"github.com/gomodule/redigo/redis"
// 	"golang.org/x/sync/semaphore"
// )

// type RedisCache struct {
// 	pool       *redis.Pool
// 	semaphore  *semaphore.Weighted
// 	redisMutex *redsync.Mutex
// 	fileMutex  *redsync.Mutex
// }

// func NewCache(host string) *Cache {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     host,
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// 	err := client.Ping().Err()

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println("\nRedis OK")
// 	return &Cache{
// 		client: client,
// 	}
// }

// var instance *RedisCache
// var once sync.Once

// func GetInstance() *RedisCache {
// 	if !config.GetBool("useRedis") {

// 		return instance
// 	}
// 	var err error
// 	once.Do(func() {
// 		instance = &RedisCache{nil, nil, nil, nil}
// 		instance.pool, err = instance.newPool()
// 		instance.semaphore = semaphore.NewWeighted(int64(config.GetInt("redisMaxConnections")))
// 		if err != nil {

// 			return
// 		}
// 	})
// 	return instance
// }

// // refresh Pool is used to refresh the redis pool
// func (cache *RedisCache) refreshPool() error {
// 	var err error
// 	cache.pool, err = cache.newPool()
// 	return err
// }

// func (cache *RedisCache) GetConnection() redis.Conn {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	return c
// }

// func (cache *RedisCache) ReleaseConnection(conn redis.Conn) {
// 	conn.Close()
// 	cache.semaphore.Release(1)
// }

// func (cache *RedisCache) newPool() (*redis.Pool, error) {
// 	var redErr error
// 	pool := redis.Pool{
// 		MaxIdle:   config.GetInt("redisMaxIdleConnections"),
// 		MaxActive: config.GetInt("redisMaxConnections"), // max number of connections
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial(config.GetString("redisCon"), config.GetString("redisAddress"))
// 			if err != nil {
// 				redErr = err

// 			}
// 			return c, err
// 		},
// 	}
// 	rs := redsync.New(redigo.NewPool(&pool))
// 	cache.redisMutex = rs.NewMutex("main-files-mutex")
// 	cache.fileMutex = rs.NewMutex("single-files-mutex")
// 	return &pool, redErr
// }

// func (cache *RedisCache) Set(key string, value []byte) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SET", key, value)
// 	if err != nil {

// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {Set} key:", key, " error:", err)
// 		}
// 		v := string(value)
// 		if len(v) > 15 {
// 			v = v[0:12] + "..."
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) GetInt(key string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("TError: ("+config.GetString("serviceLogName")+"),acquire semaphore:", err)
// 		}
// 		return -1, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	var data int64
// 	dataint, err := c.Do("GET", key)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {GetInt} key:", key, " error:", err)
// 		}
// 		return -1, err
// 	}
// 	if dataint != nil {
// 		data, err = redis.Int64(dataint, err)
// 	}
// 	return data, err
// }

// func (cache *RedisCache) SetInt(key string, value int) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)

// 	defer c.Close()
// 	_, err := c.Do("SET", key, value)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {SetInt} key:", key, " val:", value, " error:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) Increment(key string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return -1, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)

// 	defer c.Close()
// 	val, err := c.Do("INCR", key)
// 	if err != nil {

// 		return -1, err
// 	}
// 	return val.(int64), err
// }

// func (cache *RedisCache) Decrement(key string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return -1, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)

// 	defer c.Close()
// 	val, err := c.Do("DECR", key)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {Decrement} key:", key, " err:", err)
// 		}
// 		return -1, err
// 	}
// 	return val.(int64), err
// }

// func (cache *RedisCache) SetString(key string, value string) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)

// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SET", key, value)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {SetString} key:", key, " val:", value, " err:", err)
// 		}
// 		v := string(value)
// 		if len(v) > 15 {
// 			v = v[0:12] + "..."
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) Get(key string) ([]byte, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	var data []byte
// 	dataint, err := c.Do("GET", key)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {Get} key:", key, " err:", err)
// 		}
// 		return []byte{}, err
// 	}
// 	if dataint != nil {
// 		data, err = redis.Bytes(dataint, err)
// 	}
// 	return data, err
// }

// func (cache *RedisCache) GetKeys(pattern string) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()

// 	data, _ := redis.Strings(c.Do("Keys", pattern))
// 	return data
// }
// func (cache *RedisCache) GetString(key string) (string, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return "", err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	var data string
// 	dataint, err := c.Do("GET", key)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {GetString} key:", key, " err:", err)
// 		}
// 		return "", err
// 	}
// 	if dataint != nil {
// 		data, err = redis.String(dataint, err)
// 	}
// 	return data, err
// }

// func (cache *RedisCache) Del(key string) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("DEL", key)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {Del} key:", key, " err:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) Append(key string, value interface{}) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("APPEND", key, value)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {Append} key:", key, " val:", value, " err:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) SAdd(value []interface{}) error {
// 	if len(value) <= 1 {
// 		return errors.New("Not enough parameters")
// 	}
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SADD", value...)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {SAdd} key:", value, " err:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) ReadSet(key string) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("SMEMBERS", key))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", key, " err:", err)
// 		}
// 	}
// 	return data
// }
// func (cache *RedisCache) SRem(value []interface{}) error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SREM", value...)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", value, " err:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) ReadRemoveSetGPSInfoBytes(key string) []models.GPSInfo {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := c.Do("SMEMBERS", key)
// 	switch reply := data.(type) {
// 	case []interface{}:
// 		gpsArray := make([]models.GPSInfo, 0)
// 		removing := []interface{}{key}
// 		for i := range reply {
// 			if reply[i] == nil {
// 				continue
// 			}
// 			var gps models.GPSInfo
// 			gps.DecodeRedisData(reply[i].([]byte))
// 			removing = append(removing, gps.EncodeRedisData())
// 			gpsArray = append(gpsArray, gps)
// 		}
// 		c.Do("SREM", removing...)
// 		return gpsArray
// 	}

// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", key, " err:", err)
// 		}
// 	}
// 	return nil
// }

// func (cache *RedisCache) ReadGPSInfoBytes(key string) []models.GPSInfo {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := c.Do("SMEMBERS", key)
// 	switch reply := data.(type) {
// 	case []interface{}:
// 		gpsArray := make([]models.GPSInfo, 0)
// 		for i := range reply {
// 			if reply[i] == nil {
// 				continue
// 			}
// 			var gps models.GPSInfo
// 			gps.DecodeRedisData(reply[i].([]byte))
// 			gpsArray = append(gpsArray, gps)
// 		}
// 		return gpsArray
// 	}

// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", key, " err:", err)
// 		}
// 	}
// 	return nil
// }

// func (cache *RedisCache) SetISMember(key string, member string) (bool, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return false, nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	val, err := redis.Bool(c.Do("SISMEMBER", key, member))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {SetISMember} key:", key, "member:", member, " err:", err)
// 		}
// 		return false, nil
// 	}
// 	return val, err
// }

// func (cache *RedisCache) SortedSetAdd(key string, seq int32, value interface{}) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("ZADD", key, seq, value)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ZAdd} key:", key, " err:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) ReadSortedSet(key string) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSortedSet} key:", key, " err:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) RemoveSortedMessage(key string, val string) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("ZREM", key, val)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {RemoveSortedMessage} key:", key, " error:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) HashMultiSet(key string, args map[string]interface{}) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HMSET", redis.Args{key}.AddFlat(args)...)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashMultiSet} key:", key, " error:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) HashMultiSetString(key string, args map[string]string) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HMSET", redis.Args{key}.AddFlat(args)...)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashMultiSet} key:", key, " error:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) HashMultiSetInt(key string, args map[string]int) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HMSET", redis.Args{key}.AddFlat(args)...)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashMultiSetInt} key:", key, " args", args, " error:", err)
// 		}
// 	}
// }

// // HashSet first index should be key
// func (cache *RedisCache) HashSet(args []interface{}) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HSET", args...)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashSet} key:", args, " err:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) LPush(key string, args []byte) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		return
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("LPUSH", key, args)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {LPUSH} key:", args, " err:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) LTrim(key string, start int, end int) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		return
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("LTRIM", key, start, end)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {LTRIM} key:", key, " err:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) LRange(key string, start int, end int) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		return
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := redis.Strings(c.Do("LRANGE", key, start, end))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {LTRIM} key:", key, " err:", err)
// 		}
// 	}
// }

// func (cache *RedisCache) HashGet(key string, field string) string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.String(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGet} key:", key, ",Field:", field, "err:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetNoPrint(key string, field string) (string, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.String(c.Do("HGET", key, field))
// 	return data, err
// }

// func (cache *RedisCache) HashGetBytes(key string, field string) []byte {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Bytes(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGet} key:", key, ",Field:", field, "err:", err)
// 		}
// 	}
// 	return data
// }
// func (cache *RedisCache) HashDel(key string, field string) error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HDEL", key, field)
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashDel} key:", key, " err:", err)
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) HashGetAll(key string) map[string]string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.StringMap(c.Do("HGETALL", key))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetAll} key:", key, " err:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetInt(key string, field string) int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Int(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt} key:", key, " error:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetBool(key string, field string) bool {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Bool(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetBool} key:", key, " error:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetInt64(key string, field string) int64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Int64(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt64} key:", key, " err:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetFloat64(key string, field string) float64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Float64(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt64} key:", key, " err:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetInt64WithError(key string, field string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Int64(c.Do("HGET", key, field))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt64} key:", key, " err:", err)
// 		}
// 	}
// 	return data, err
// }

// func (cache *RedisCache) HashMGet(field []interface{}) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("HMGET", field...))
// 	if err != nil {
// 		if config.GetBool("printRedis") {
// 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {HMGET} key:", field, " err:", err)
// 		}
// 	}
// 	return data
// }

// func (cache *RedisCache) HashMGetInts(field []interface{}) []int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Ints(c.Do("HMGET", field...))
// 	if err != nil {

// 	}
// 	return data
// }

// func (cache *RedisCache) HashIncrementBy(key string, field string, val int) int64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	value, err := redis.Int64(c.Do("HINCRBY", key, field, val))
// 	if err != nil {

// 	}
// 	return value
// }
// func (cache *RedisCache) HashIncrementByFloat(key string, field string, val float64) int64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	value, err := redis.Int64(c.Do("HINCRBYFLOAT", key, field, val))
// 	if err != nil {

// 	}
// 	return value
// }

// func (cache *RedisCache) Exists(key string) bool {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Bool(c.Do("EXISTS", key))
// 	if err != nil {

// 	}
// 	return data
// }

// func (cache *RedisCache) GetMinSortedSet(key string) int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("ZRANGEBYSCORE", key, "-inf", "+inf", "WITHSCORES", "LIMIT", 0, 1))
// 	if err != nil {

// 	}
// 	if len(data) != 0 {
// 		val, _ := strconv.Atoi(data[1])
// 		return val
// 	}
// 	return 0
// }

// func (cache *RedisCache) GetMaxSortedSet(key string) int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("ZRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "LIMIT", 0, 1))
// 	if err != nil {

// 	}
// 	if len(data) != 0 {
// 		val, _ := strconv.Atoi(data[1])
// 		return val
// 	}
// 	return 0
// }

// func (cache *RedisCache) GetMutexLock() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if err := cache.redisMutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// func (cache *RedisCache) UnlockMutex() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if _, err := cache.redisMutex.Unlock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// func (cache *RedisCache) GetFileMutexLock() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if err := cache.fileMutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// func (cache *RedisCache) UnlockFileMutex() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if _, err := cache.fileMutex.Unlock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// //////

// package cachedata

// import (
// 	"context"
// 	"errors"
// 	"log"
// 	"strconv"
// 	"sync"

// 	"github.com/go-redsync/redsync/v4"
// 	"github.com/go-redsync/redsync/v4/redis/redigo"
// 	"github.com/gomodule/redigo/redis"
// 	"github.com/websays-intelligence/wutils/wcache"
// 	"golang.org/x/sync/semaphore"
// )

// type RedisCache struct {
// 	pool       *redis.Pool
// 	semaphore  *semaphore.Weighted
// 	redisMutex *redsync.Mutex
// 	fileMutex  *redsync.Mutex
// }

// var instance *RedisCache
// var once sync.Once

// func GetInstance() *RedisCache {

// 	var err error
// 	once.Do(func() {
// 		instance = &RedisCache{nil, nil, nil, nil}
// 		// cachestring := wcache.GetRedisHostFromEnv()
// 		instance.pool, err = instance.newPool()
// 		instance.semaphore = semaphore.NewWeighted(1000)
// 		if err != nil {

// 			return
// 		}
// 	})
// 	return instance
// }

// // refresh Pool is used to refresh the redis pool
// func (cache *RedisCache) refreshPool() error {
// 	var err error
// 	cache.pool, err = cache.newPool()
// 	return err
// }

// func (cache *RedisCache) GetConnection() redis.Conn {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	return c
// }

// func (cache *RedisCache) ReleaseConnection(conn redis.Conn) {
// 	conn.Close()
// 	cache.semaphore.Release(1)
// }

// func (cache *RedisCache) newPool() (*redis.Pool, error) {
// 	var redErr error
// 	redisString := wcache.GetRedisHostFromEnv()
// 	log.Println("RedisString: ", redisString)
// 	pool := redis.Pool{

// 		MaxIdle:   80,
// 		MaxActive: 1000, // max number of connections
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial("tcp", redisString)
// 			if err != nil {
// 				redErr = err

// 			}
// 			return c, err
// 		},
// 	}
// 	rs := redsync.New(redigo.NewPool(&pool))
// 	cache.redisMutex = rs.NewMutex("main-files-mutex")
// 	cache.fileMutex = rs.NewMutex("single-files-mutex")
// 	return &pool, redErr
// }

// func (cache *RedisCache) Set(key string, value []byte) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SET", key, value)
// 	if err != nil {

// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {Set} key:", key, " error:", err)
// 		// }
// 		// v := string(value)
// 		// if len(v) > 15 {
// 		// 	v = v[0:12] + "..."
// 		// }
// 	}
// 	return err
// }

// func (cache *RedisCache) GetInt(key string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("TError: ("+config.GetString("serviceLogName")+"),acquire semaphore:", err)
// 		// }
// 		return -1, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	var data int64
// 	dataint, err := c.Do("GET", key)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {GetInt} key:", key, " error:", err)
// 		// }
// 		return -1, err
// 	}
// 	if dataint != nil {
// 		data, err = redis.Int64(dataint, err)
// 	}
// 	return data, err
// }

// func (cache *RedisCache) SetInt(key string, value int) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)

// 	defer c.Close()
// 	_, err := c.Do("SET", key, value)
// 	// if err != nil {
// 	// if config.GetBool("printRedis") {
// 	// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {SetInt} key:", key, " val:", value, " error:", err)
// 	// }
// 	// }
// 	return err
// }

// func (cache *RedisCache) Increment(key string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return -1, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)

// 	defer c.Close()
// 	val, err := c.Do("INCR", key)
// 	if err != nil {

// 		return -1, err
// 	}
// 	return val.(int64), err
// }

// func (cache *RedisCache) Decrement(key string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return -1, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)

// 	defer c.Close()
// 	val, err := c.Do("DECR", key)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {Decrement} key:", key, " err:", err)
// 		// }
// 		return -1, err
// 	}
// 	return val.(int64), err
// }

// func (cache *RedisCache) SetString(key string, value string) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)

// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SET", key, value)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {SetString} key:", key, " val:", value, " err:", err)
// 		// }
// 		v := string(value)
// 		if len(v) > 15 {
// 			v = v[0:12] + "..."
// 		}
// 	}
// 	return err
// }

// func (cache *RedisCache) Get(key string) ([]byte, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil, err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	var data []byte
// 	dataint, err := c.Do("GET", key)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {Get} key:", key, " err:", err)
// 		// }
// 		return []byte{}, err
// 	}
// 	if dataint != nil {
// 		data, err = redis.Bytes(dataint, err)
// 	}
// 	return data, err
// }

// func (cache *RedisCache) GetKeys(pattern string) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()

// 	data, _ := redis.Strings(c.Do("Keys", pattern))
// 	return data
// }
// func (cache *RedisCache) GetString(key string) (string, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return "", err
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	var data string
// 	dataint, err := c.Do("GET", key)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {GetString} key:", key, " err:", err)
// 		// }
// 		return "", err
// 	}
// 	if dataint != nil {
// 		data, err = redis.String(dataint, err)
// 	}
// 	return data, err
// }

// func (cache *RedisCache) Del(key string) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("DEL", key)
// 	// if err != nil {
// 	// if config.GetBool("printRedis") {
// 	// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {Del} key:", key, " err:", err)
// 	// }
// 	// }
// 	return err
// }

// func (cache *RedisCache) Append(key string, value interface{}) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("APPEND", key, value)
// 	// if err != nil {
// 	// if config.GetBool("printRedis") {
// 	// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {Append} key:", key, " val:", value, " err:", err)
// 	// }
// 	// }
// 	return err
// }

// func (cache *RedisCache) SAdd(value []interface{}) error {
// 	if len(value) <= 1 {
// 		return errors.New("not enough parameters")
// 	}
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SADD", value...)
// 	// if err != nil {
// 	// if config.GetBool("printRedis") {
// 	// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {SAdd} key:", value, " err:", err)
// 	// }
// 	// }
// 	return err
// }

// func (cache *RedisCache) ReadSet(key string) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("SMEMBERS", key))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", key, " err:", err)
// 		// }
// 		// return
// 	}
// 	return data
// }
// func (cache *RedisCache) SRem(value []interface{}) error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("SREM", value...)
// 	// if err != nil {
// 	// if config.GetBool("printRedis") {
// 	// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", value, " err:", err)
// 	// }
// 	// }
// 	return err
// }

// // func (cache *RedisCache) ReadRemoveSetGPSInfoBytes(key string) []models.GPSInfo {
// // 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// // 		return nil
// // 	}
// // 	c := cache.pool.Get()
// // 	defer cache.semaphore.Release(1)
// // 	defer c.Close()
// // 	data, err := c.Do("SMEMBERS", key)
// // 	switch reply := data.(type) {
// // 	case []interface{}:
// // 		gpsArray := make([]models.GPSInfo, 0)
// // 		removing := []interface{}{key}
// // 		for i := range reply {
// // 			if reply[i] == nil {
// // 				continue
// // 			}
// // 			var gps models.GPSInfo
// // 			gps.DecodeRedisData(reply[i].([]byte))
// // 			removing = append(removing, gps.EncodeRedisData())
// // 			gpsArray = append(gpsArray, gps)
// // 		}
// // 		c.Do("SREM", removing...)
// // 		return gpsArray
// // 	}

// // 	if err != nil {
// // 		if config.GetBool("printRedis") {
// // 			log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", key, " err:", err)
// // 		}
// // 	}
// // 	return nil
// // }

// // func (cache *RedisCache) ReadGPSInfoBytes(key string) []models.GPSInfo {
// // 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// // 		return nil
// // 	}
// // 	c := cache.pool.Get()
// // 	defer cache.semaphore.Release(1)
// // 	defer c.Close()
// // 	data, err := c.Do("SMEMBERS", key)
// // 	switch reply := data.(type) {
// // 	case []interface{}:
// // 		gpsArray := make([]models.GPSInfo, 0)
// // 		for i := range reply {
// // 			if reply[i] == nil {
// // 				continue
// // 			}
// // 			var gps models.GPSInfo
// // 			gps.DecodeRedisData(reply[i].([]byte))
// // 			gpsArray = append(gpsArray, gps)
// // 		}
// // 		return gpsArray
// // 	}

// // 	if err != nil {
// // 		// if config.GetBool("printRedis") {
// // 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSet} key:", key, " err:", err)
// // 		// }
// // 	}
// // 	return nil
// // }

// func (cache *RedisCache) SetISMember(key string, member string) (bool, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return false, nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	val, err := redis.Bool(c.Do("SISMEMBER", key, member))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {SetISMember} key:", key, "member:", member, " err:", err)
// 		// }
// 		return false, nil
// 	}
// 	return val, err
// }

// func (cache *RedisCache) SortedSetAdd(key string, seq int32, value interface{}) error {
// 	cache.semaphore.Acquire(context.TODO(), 1)
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("ZADD", key, seq, value)
// 	// if err != nil {
// 	// if config.GetBool("printRedis") {
// 	// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {ZAdd} key:", key, " err:", err)
// 	// }
// 	// }
// 	return err
// }

// func (cache *RedisCache) ReadSortedSet(key string) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 		return nil
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {ReadSortedSet} key:", key, " err:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) RemoveSortedMessage(key string, val string) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("ZREM", key, val)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {RemoveSortedMessage} key:", key, " error:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) HashMultiSet(key string, args map[string]interface{}) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HMSET", redis.Args{key}.AddFlat(args)...)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashMultiSet} key:", key, " error:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) HashMultiSetString(key string, args map[string]string) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HMSET", redis.Args{key}.AddFlat(args)...)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashMultiSet} key:", key, " error:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) HashMultiSetInt(key string, args map[string]int) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HMSET", redis.Args{key}.AddFlat(args)...)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashMultiSetInt} key:", key, " args", args, " error:", err)
// 		// }
// 	}
// }

// // HashSet first index should be key
// func (cache *RedisCache) HashSet(args []interface{}) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HSET", args...)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashSet} key:", args, " err:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) LPush(key string, args []byte) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		return
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("LPUSH", key, args)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {LPUSH} key:", args, " err:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) LTrim(key string, start int, end int) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		return
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("LTRIM", key, start, end)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {LTRIM} key:", key, " err:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) LRange(key string, start int, end int) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {
// 		return
// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := redis.Strings(c.Do("LRANGE", key, start, end))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {LTRIM} key:", key, " err:", err)
// 		// }
// 	}
// }

// func (cache *RedisCache) HashGet(key string, field string) string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.String(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGet} key:", key, ",Field:", field, "err:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetNoPrint(key string, field string) (string, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.String(c.Do("HGET", key, field))
// 	return data, err
// }

// func (cache *RedisCache) HashGetBytes(key string, field string) []byte {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Bytes(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGet} key:", key, ",Field:", field, "err:", err)
// 		// }
// 	}
// 	return data
// }
// func (cache *RedisCache) HashDel(key string, field string) error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	_, err := c.Do("HDEL", key, field)
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashDel} key:", key, " err:", err)
// 		// }
// 	}
// 	return err
// }

// func (cache *RedisCache) HashGetAll(key string) map[string]string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.StringMap(c.Do("HGETALL", key))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetAll} key:", key, " err:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetInt(key string, field string) int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Int(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt} key:", key, " error:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetBool(key string, field string) bool {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Bool(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetBool} key:", key, " error:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetInt64(key string, field string) int64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Int64(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt64} key:", key, " err:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetFloat64(key string, field string) float64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Float64(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt64} key:", key, " err:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashGetInt64WithError(key string, field string) (int64, error) {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Int64(c.Do("HGET", key, field))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HashGetInt64} key:", key, " err:", err)
// 		// }
// 	}
// 	return data, err
// }

// func (cache *RedisCache) HashMGet(field []interface{}) []string {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("HMGET", field...))
// 	if err != nil {
// 		// if config.GetBool("printRedis") {
// 		// 	log.Println("RError: ("+config.GetString("serviceLogName")+"), {HMGET} key:", field, " err:", err)
// 		// }
// 	}
// 	return data
// }

// func (cache *RedisCache) HashMGetInts(field []interface{}) []int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Ints(c.Do("HMGET", field...))
// 	if err != nil {

// 	}
// 	return data
// }

// func (cache *RedisCache) HashIncrementBy(key string, field string, val int) int64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	value, err := redis.Int64(c.Do("HINCRBY", key, field, val))
// 	if err != nil {

// 	}
// 	return value
// }
// func (cache *RedisCache) HashIncrementByFloat(key string, field string, val float64) int64 {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	value, err := redis.Int64(c.Do("HINCRBYFLOAT", key, field, val))
// 	if err != nil {

// 	}
// 	return value
// }

// func (cache *RedisCache) Exists(key string) bool {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Bool(c.Do("EXISTS", key))
// 	if err != nil {

// 	}
// 	return data
// }

// func (cache *RedisCache) GetMinSortedSet(key string) int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("ZRANGEBYSCORE", key, "-inf", "+inf", "WITHSCORES", "LIMIT", 0, 1))
// 	if err != nil {

// 	}
// 	if len(data) != 0 {
// 		val, _ := strconv.Atoi(data[1])
// 		return val
// 	}
// 	return 0
// }

// func (cache *RedisCache) GetMaxSortedSet(key string) int {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	c := cache.pool.Get()
// 	defer cache.semaphore.Release(1)
// 	defer c.Close()
// 	data, err := redis.Strings(c.Do("ZRANGEBYSCORE", key, "+inf", "-inf", "WITHSCORES", "LIMIT", 0, 1))
// 	if err != nil {

// 	}
// 	if len(data) != 0 {
// 		val, _ := strconv.Atoi(data[1])
// 		return val
// 	}
// 	return 0
// }

// func (cache *RedisCache) GetMutexLock() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if err := cache.redisMutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// func (cache *RedisCache) UnlockMutex() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if _, err := cache.redisMutex.Unlock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// func (cache *RedisCache) GetFileMutexLock() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if err := cache.fileMutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }

// func (cache *RedisCache) UnlockFileMutex() error {
// 	if err := cache.semaphore.Acquire(context.TODO(), 1); err != nil {

// 	}
// 	if _, err := cache.fileMutex.Unlock(); err != nil {
// 		return err
// 	}
// 	defer cache.semaphore.Release(1)
// 	return nil
// }
