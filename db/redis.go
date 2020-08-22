package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis/v8"
)

// global vars containing SHA1 digests of Lua scipts once they have been loaded
// into redis.
var (
	saveLua               = ""
	getAndIncrLua         = ""
	deleteShortAndLongLua = ""
)

// RedisDB is the DB implementation using Redis.
type RedisDB struct {
	rdb *redis.Client
}

// NewRedisDB returns an instance of the RedisDB ready to use.
func NewRedisDB(ctx context.Context, addr string) (*RedisDB, error) {
	rs := &RedisDB{
		rdb: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
	// Try to ping the service.
	err := rs.rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	err = rs.loadScripts(ctx)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

// Save a URL and its shortened version.
func (db *RedisDB) Save(ctx context.Context, url, hash string) error {
	// run pre-loaded script
	_, err := db.rdb.EvalSha(
		ctx,
		saveLua,
		[]string{url, hash}, // KEYS
		[]string{hash, url}, // ARGV
	).Result()
	return err
}

// Find returns the value matching the given key.
func (db *RedisDB) Find(ctx context.Context, key string) (string, error) {
	return db.rdb.Get(ctx, key).Result()
}

// FindAndIncr returns the value matching the given key and increments the
// counter.
func (db *RedisDB) FindAndIncr(ctx context.Context, key string) (string, error) {
	// run pre-loaded script
	hash, err := db.rdb.EvalSha(
		ctx,
		getAndIncrLua,
		[]string{key}, // KEYS
	).Result()
	return hash.(string), err
}

// Delete removes both the URL and its shortened version.
func (db *RedisDB) Delete(ctx context.Context, key string) error {
	// run pre-loaded script
	_, err := db.rdb.EvalSha(
		ctx,
		deleteShortAndLongLua,
		[]string{key}, // KEYS
	).Result()
	return err
}

// Count return the redirections count value for the given URL.
func (db *RedisDB) Count(ctx context.Context, url string) (int, error) {
	return db.rdb.Get(ctx, fmt.Sprintf("count:%s", url)).Int()
}

// Incr increments the redirections count value for the given URL.
func (db *RedisDB) Incr(ctx context.Context, url string) error {
	_, err := db.rdb.Incr(ctx, fmt.Sprintf("count:%s", url)).Result()
	return err
}

func (r *RedisDB) loadScripts(ctx context.Context) error {
	saveStr, err := readLuaScript("save.lua")
	if err != nil {
		return err
	}

	saveLua, err = r.rdb.ScriptLoad(ctx, saveStr).Result()
	if err != nil {
		return err
	}

	getAndIncrStr, err := readLuaScript("getAndIncr.lua")
	if err != nil {
		return err
	}

	getAndIncrLua, err = r.rdb.ScriptLoad(ctx, getAndIncrStr).Result()
	if err != nil {
		return err
	}

	deleteShortAndLongStr, err := readLuaScript("deleteShortAndLong.lua")
	if err != nil {
		return err
	}

	deleteShortAndLongLua, err = r.rdb.ScriptLoad(ctx, deleteShortAndLongStr).Result()
	if err != nil {
		return err
	}

	return nil
}

func readLuaScript(fileName string) (string, error) {
	file, err := os.Open(fmt.Sprintf("../scripts/%s", fileName))
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(b), err
}
