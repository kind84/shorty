package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis/v7"
)

// global vars containing SHA1 digests of Lua scipts once they have been loaded
// into redis.
var (
	saveAndCountLua = ""
)

type RedisDB struct {
	rdb *redis.Client
}

func NewRedisDB(addr string) (*RedisDB, error) {
	rs := &RedisDB{
		rdb: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
	// Try to ping the service.
	err := rs.rdb.Ping().Err()
	if err != nil {
		return nil, err
	}

	// err = rs.loadScripts()
	// if err != nil {
	// 	return nil, err
	// }

	return rs, nil
}

func (db *RedisDB) Save(ctx context.Context, url, hash string) error {
	return nil
}

func (db *RedisDB) Find(ctx context.Context, hash string) (string, error) {
	return "", nil
}

func (db *RedisDB) Delete(context.Context, string) error {
	return nil
}

func (r *RedisDB) loadScripts() error {
	saveAndCountStr, err := readLuaScript("saveAndCount.lua")
	if err != nil {
		return err
	}

	saveAndCountLua, err = r.rdb.ScriptLoad(saveAndCountStr).Result()
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
