-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class comp_redis
---@field dial fun(opt: table):redis_conn

---@class redis_conn
---@field set fun(self: redis_conn, k: string, v: string, timeout?: timeDuration): err
---@field setnx fun(self: redis_conn, k: string, v: string, timeout?: timeDuration): err
---@field get fun(self: redis_conn, k: string): string
---@field del fun(self: redis_conn, keys: ...): err
---@field eval fun(self: redis_conn, script: string, keys: string[], args: string[]): table
---@field hset fun(self: redis_conn, key: string, fields: table<string, string>)
---@field hdel fun(self: redis_conn, key: string, fields: string[])
---@field hget fun(self: redis_conn, key: string, field: string): string, err
---@field scan fun(self: redis_conn, match: string, count: integer)
---@field sscan fun(self: redis_conn, key: string, match: string, count: integer)
---@field lrange fun(self: redis_conn, key: string, start: integer, end: integer)
---@field lpush fun(self: redis_conn, key: string, values: string[])
---@field rpush fun(self: redis_conn, key: string, values: string[])
---@field lpop fun(self: redis_conn, key: string)
---@field rpop fun(self: redis_conn, key: string)
---@field incr fun(self: redis_conn, key: string)
---@field decr fun(self: redis_conn, key: string)
---@field ping fun(self: redis_conn)
---@field close fun(self: redis_conn)
---@field exec fun(self: redis_conn, cmd: string): err
