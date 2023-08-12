-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

ffi.library("rabbitmq", {
    Dial = { "ptr", { "str" } },
    Consume = { "void", { "ptr", "str" } },
    QueueDeclare = { "str", { "ptr", "str" } },
    ExchangeDeclare = { "str", { "ptr", "str" } },
    QueueBind = { "str", { "ptr", "str" } },
    Publish = { "str", { "ptr", "str" } },
})

local rabbitmq_conn = {}

function rabbitmq_conn:queue_declare(opt)
    ffi.lib.rabbitmq.QueueDeclare(self.conn, json.encode(opt))
    return self
end

function rabbitmq_conn:exchange_declare(opt)
    ffi.lib.rabbitmq.ExchangeDeclare(self.conn, json.encode(opt))
    return self
end

function rabbitmq_conn:queue_bind(opt)
    ffi.lib.rabbitmq.QueueBind(self.conn, json.encode(opt))
    return self
end

function rabbitmq_conn:publish(opt)
    return ffi.lib.rabbitmq.Publish(self.conn, json.encode(opt))
end

function rabbitmq_conn:consume(opt)
    if opt.callback == nil then
        return
    end
    local port = random.port()
    opt.url = string.format("http://localhost:%d", port)
    local cb = opt.callback
    opt.callback = nil
    go(function()
        local s = mock.new()
        s:post("/", function(ctx)
            local d, _ = io.ReadAll(ctx.Request.Body)
            local res = ""
            for i = 1, #d, 1 do
                res = res .. string.char(d[i])
            end
            cb(res)
        end)
        yassert(s:run(port))
    end)
    ffi.lib.rabbitmq.Consume(self.conn, json.encode(opt))
end

return {
    ---@return rabbitmq_conn
    dial = function(dataSourceName)
        local obj = {
            conn = ffi.lib.rabbitmq.Dial(dataSourceName)
        }
        setmetatable(obj, { __index = rabbitmq_conn })
        return obj
    end
}
