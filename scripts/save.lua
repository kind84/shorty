redis.call("set", KEYS[1], ARGV[1])
redis.call("set", KEYS[2], ARGV[2])
return true
