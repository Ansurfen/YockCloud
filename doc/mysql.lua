-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class comp_mysql
---@field dial fun(dataSourceName: string): mysql_conn

---@class mysql_conn
---@field exec fun(self: mysql_conn, sql: string): err
---@field query fun(self: mysql_conn, sql: string): table
---@field ping fun(self: mysql_conn)
---@field close fun(self: mysql_conn)
