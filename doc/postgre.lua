-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class comp_postgre
---@field dial fun(dataSourceName: string): postgre_conn

---@class postgre_conn
---@field exec fun(self: postgre_conn, sql: string): err
---@field query fun(self: postgre_conn, sql: string): table
---@field ping fun(self: postgre_conn)
---@field close fun(self: postgre_conn)
