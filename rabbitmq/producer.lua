-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@module 'rabbitmq'
local rabbitmq = import("./index")

local conn = rabbitmq.dial("amqp://guest:guest@192.168.127.128:5673/")
conn:queue_declare({
    name = "queue_publisher",
    durable = true,
    auto_delete = false,
    exclusive = false,
    nowait = false
}):exchange_declare({
    name = "exchange_publisher",
    kind = "topic",
    durable = true,
    auto_delete = false,
    internal = false,
    nowait = false,
}):queue_bind({
    name = "queue_publisher",
    key = "key1",
    exchange = "exchange_publisher",
    nowait = true
})
for i = 1, 10, 1 do
    conn:publish({
        exchange = "exchange_publisher",
        key = "key1",
        mandatory = false,
        immediate = false,
        msg = {
            content_type = "text/plain",
            body = tostring(i)
        }
    })
    time.Sleep(1 * time.Second)
end
