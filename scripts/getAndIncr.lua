local short = redis.call("get", KEYS[1])
if short != "" then
	redis.call("incr", KEYS[1])
end
return short 
