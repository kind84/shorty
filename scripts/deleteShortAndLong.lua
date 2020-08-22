local short = redis.call("get", KEYS[1])
redis.call("del", KEYS[1])
if short ~= false then 
	redis.call("del", short)
	redis.call("del", "count:"..short)
end
return true
