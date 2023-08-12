-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class etcd_option
---@field dial-timeout timeDuration
---@field endpoints string[]

---@class comp_etcd
---@field dial fun(opt: etcd_option): etcd_conn

---@class etcd_watch_option
---@field key string
---@field callback fun(res: table)
---@field url? string

---@class etcd_conn
---@field put fun(self: etcd_conn, k: string, v: string): table
---@field get fun(self: etcd_conn, k: string): table
---@field watch fun(self: etcd_conn, opt: etcd_watch_option)
