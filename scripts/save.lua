if redis.call("set", KEYS[1], ARGV[1]) == 1 then
	return redis.call("set", KEYS[2], ARGV[2])
end
return false
