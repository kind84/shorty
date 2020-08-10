local short = redis.call("get", KEYS[1])
redis.call("del", KEYS[1])
redis.call("del", short)
redis.call("del", "count:"..KEYS[1])
return true
