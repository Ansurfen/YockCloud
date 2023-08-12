-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@type comp_rabbitmq
local rabbitmq = import("./index")

local conn = rabbitmq.dial("amqp://guest:guest@192.168.127.128:5673/")
local ch = channel.make()
go(function()
    conn:queue_declare({
        name = "queue_publisher",
        durable = true,
        auto_delete = false,
        exclusive = false,
        nowait = false,
    }):consume({
        queue = "queue_publisher",
        consumer = "",
        auto_ack = true,
        exclusive = true,
        nolocal = false,
        nowait = true,
        callback = function(body)
            ch:send(body)
        end
    })
end)
go(function()
    while true do
        channel.select({
            "|<-", ch, function(ok, v)
            print(ok, v)
        end })
    end
end)
wait("blocked")
