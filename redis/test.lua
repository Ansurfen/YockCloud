-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@type comp_redis
local redis = import("./index")

local conn = redis.dial({ addr = "192.168.127.128:6379" })
-- conn:set("a", "6")
-- table.dump(conn:eval({
--     script = "redis.call('SET',KEYS[1],ARGV[1]);redis.call('EXPIRE',KEYS[1],ARGV[2]);return 1;",
--     keys = { "1", "key1" },
--     args = { 10, 60 },
-- }))
-- print(conn:get("1"))
-- conn:rpush("aa", { "a", "b" })
-- conn:lrange("aa", 0, 10)
-- conn:scan("a", 1)

-- print(conn:hset("counter", { a = "c" }))
-- -- conn:hdel("counter", { "a" })
-- print(conn:del("counter"))
-- print(conn:hget("counter", "a"))
print(conn:exec("del a"))
print(conn:exec("get a"))
