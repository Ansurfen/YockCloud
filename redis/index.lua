-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global, undefined-field
ffi.library("redis", {
    Dial = { "ptr", { "str" } },
    Set = { "str", { "ptr", "str", "str", "double" } },
    SetNX = { "str", { "ptr", "str", "str", "double" } },
    Get = { "str", { "ptr", "str" } },
    Del = { "str", { "ptr", "str" } },
    Eval = { "str", { "ptr", "str" } },
    HSet = { "str", { "ptr", "str" } },
    HGet = { "str", { "ptr", "str", "str" } },
    HDel = { "str", { "ptr", "str" } },
    LPush = { "str", { "ptr", "str" } },
    RPush = { "str", { "ptr", "str" } },
    LRange = { "str", { "ptr", "str" } },
    LPop = { "str", { "ptr", "str" } },
    RPop = { "str", { "ptr", "str" } },
    Incr = { "str", { "ptr", "str" } },
    Decr = { "str", { "ptr", "str" } },
    Scan = { "str", { "ptr", "str", "int64" } },
    SScan = { "str", { "ptr", "str", "str", "int64" } },
    Close = { "str", { "ptr", "str" } },
    Ping = { "str", { "ptr", "str" } },
    Do = { "str", { "ptr", "str" } }
})

local redis_conn = {}

function redis_conn:exec(cmd)
    local res = json.decode(ffi.lib.redis.Do(self.conn, json.encode(strings.Fields(cmd))))
    return res["data"], res["err"]
end

---@param k string
---@param v string
---@param timeout timeDuration
function redis_conn:set(k, v, timeout)
    if timeout == nil then
        return ffi.lib.redis.Set(self.conn, k, v, 0)
    end
    return ffi.lib.redis.Set(self.conn, k, v, timeout)
end

---@param k string
---@param v string
---@param timeout timeDuration
function redis_conn:setnx(k, v, timeout)
    if timeout == nil then
        return ffi.lib.redis.SetNX(self.conn, k, v, 0)
    end
    return ffi.lib.redis.Set(self.conn, k, v, timeout)
end

---@param k string
---@return string
function redis_conn:get(k)
    return ffi.lib.redis.Get(self.conn, k)
end

---@param k string
---@return string
function redis_conn:del(...)
    return ffi.lib.redis.Del(self.conn, json.encode({ ... }))
end

---@return err
function redis_conn:ping()
    return ffi.lib.redis.Ping(self.conn)
end

---@return err
function redis_conn:close()
    return ffi.lib.redis.Close(self.conn)
end

function redis_conn:eval(script, keys, args)
    local res = ffi.lib.redis.Eval(self.conn, json.encode({
        script = script,
        keys = keys,
        args = args
    }))
    return json.decode(res)
end

---@param k string
---@param fields table<string, string>
function redis_conn:hset(k, fields)
    local tmp = {}
    for key, value in pairs(fields) do
        table.insert(tmp, key)
        table.insert(tmp, value)
    end
    return ffi.lib.redis.HSet(self.conn, json.encode({
        key = k,
        field = tmp
    }))
end

function redis_conn:hget(key, field)
    local res = json.decode(ffi.lib.redis.HGet(self.conn, key, field))
    return res["data"], res["err"]
end

function redis_conn:hdel(key, fields)
    local tmp = {}
    for _, v in ipairs(fields) do
        table.insert(tmp, v)
    end
    return ffi.lib.redis.HDel(self.conn, json.encode({
        key = key,
        field = tmp
    }))
end

function redis_conn:lpush(k, vs)
    return ffi.lib.redis.LPush(self.conn, json.encode({
        key = k,
        values = vs
    }))
end

function redis_conn:rpush(k, vs)
    return ffi.lib.redis.RPush(self.conn, json.encode({
        key = k,
        values = vs
    }))
end

function redis_conn:lrange(k, s, e)
    local res = ffi.lib.redis.LRange(self.conn, json.encode({
        key = k,
        start = s,
        End = e
    }))
    return json.decode(res)
end

function redis_conn:scan(match, count)
    local res = json.decode(ffi.lib.redis.Scan(self.conn, match, count))
    return res
end

function redis_conn:sscan(key, match, count)
    local res = ffi.lib.redis.SScan(self.conn, key, match, count)
    return json.decode(res)
end

---@param k string
---@return string
function redis_conn:lpop(k)
    return ffi.lib.redis.LPop(self.conn, k)
end

---@param k string
---@return string
function redis_conn:rpop(k)
    return ffi.lib.redis.RPop(self.conn, k)
end

---@param k string
---@return string
function redis_conn:incr(k)
    return ffi.lib.redis.Incr(self.conn, k)
end

---@param k string
---@return string
function redis_conn:decr(k)
    return ffi.lib.redis.Decr(self.conn, k)
end

return {
    ---@param opt table
    ---@return redis_conn
    dial = function(opt)
        local obj = {
            conn = ffi.lib.redis.Dial(json.encode(opt))
        }
        setmetatable(obj, { __index = redis_conn })
        return obj
    end
}
