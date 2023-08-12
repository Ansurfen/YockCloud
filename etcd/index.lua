-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global, undefined-field
ffi.library("etcd", {
    Dial = { "ptr", { "str" } },
    Watch = { "void", { "ptr", "str" } },
    Get = { "str", { "ptr", "str" } },
    Put = { "str", { "ptr", "str", "str" } },
    -- Query = { "str", { "ptr", "str" } },
    -- Exec = { "str", { "ptr", "str" } },
    Close = { "str", { "ptr" } },
    -- Ping = { "str", { "ptr" } },
})

local etcd_conn = {}

function etcd_conn:close()
    return ffi.lib.etcd.Close(self.conn)
end

-- function etcd_conn:ping()
--     return ffi.lib.etcd.Ping(self.conn)
-- end

function etcd_conn:watch(opt)
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
            cb(json.decode(res))
        end)
        yassert(s:run(port))
    end)
    ffi.lib.etcd.Watch(self.conn, json.encode(opt))
end

function etcd_conn:put(k, v)
    return json.decode(ffi.lib.etcd.Put(self.conn, k, v))
end

function etcd_conn:get(k)
    return json.decode(ffi.lib.etcd.Get(self.conn, k))
end

return {
    dial = function(conf)
        local obj = {
            conn = ffi.lib.etcd.Dial(json.encode(conf))
        }
        setmetatable(obj, { __index = etcd_conn })
        return obj
    end,
}
