-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class comp_rabbitmq
---@field dial fun(datasourceName: string): rabbitmq_conn

---@class rabbitmq_conn
---@field queue_declare fun(self: rabbitmq_conn, opt: queue_declare_opt): rabbitmq_conn
---@field exchange_declare fun(self: rabbitmq_conn, opt: exchange_declare_opt): rabbitmq_conn
---@field queue_bind fun(self: rabbitmq_conn, opt: queue_bind_opt): rabbitmq_conn
---@field publish fun(self: rabbitmq_conn, opt: publish_opt): err
---@field consume fun(self: rabbitmq_conn, opt: consume_opt)

---@class queue_declare_opt
---@field name string
---@field durable boolean
---@field auto_delete boolean
---@field exclusive boolean
---@field nowait boolean

---@class exchange_declare_opt
---@field name string
---@field kind string
---@field durable boolean
---@field auto_delete boolean
---@field internal boolean
---@field nowait boolean

---@class queue_bind_opt
---@field name string
---@field key string
---@field exchange string
---@field nowait boolean

---@class publishing
---@field content_type string
---@field body string

---@class publish_opt
---@field exchange string
---@field key string
---@field mandatory boolean
---@field immediate boolean
---@field msg publishing

---@class consume_opt
---@field auto_ack boolean
---@field queue string
---@field consumer string
---@field nolocal boolean
---@field exclusive boolean
---@field nowait boolean
---@field callback fun(body: string)
