if redis.call("set", KEYS[1], ARGV[1]) == 1 then
	if redis.call("set", KEYS[2], ARGV[2]) == 1 then
		return redis.call("incr", KEYS[3])
	end
end
return false
