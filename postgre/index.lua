-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global, undefined-field
ffi.library("postgre", {
    Dial = { "ptr", { "str" } },
    Query = { "str", { "ptr", "str" } },
    Exec = { "str", { "ptr", "str" } },
    Close = { "str", { "ptr" } },
    Ping = { "str", { "ptr" } },
})

local postgre_conn = {}

function postgre_conn:exec(sql)
    local res = json.decode(ffi.lib.postgre.Exec(self.conn, sql))
    return res["msg"], res["err"]
end

function postgre_conn:query(sql)
    local res = json.decode(ffi.lib.postgre.Query(self.conn, sql))
    return res["msg"], res["err"]
end

function postgre_conn:close()
    return ffi.lib.postgre.Close(self.conn)
end

function postgre_conn:ping()
    return ffi.lib.postgre.Ping(self.conn)
end

return {
    dial = function(dataSourceName)
        local obj = {
            conn = ffi.lib.postgre.Dial(dataSourceName)
        }
        setmetatable(obj, { __index = postgre_conn })
        return obj
    end,
}
