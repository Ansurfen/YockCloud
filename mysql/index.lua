-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global, undefined-field
ffi.library("mysql", {
    Dial = { "ptr", { "str" } },
    Query = { "str", { "ptr", "str" } },
    Exec = { "str", { "ptr", "str" } },
    Close = { "str", { "ptr" } },
    Ping = { "str", { "ptr" } },
})

local mysql_conn = {}

function mysql_conn:exec(sql)
    local err = ffi.lib.mysql.Exec(self.conn, sql)
    if err == nil or #err == 0 then
        return nil
    end
    return err
end

function mysql_conn:query(sql)
    return json.decode(ffi.lib.mysql.Query(self.conn, sql))
end

function mysql_conn:close()
    local err = ffi.lib.mysql.Close(self.conn)
    if err == nil or #err == 0 then
        return nil
    end
    return err
end

function mysql_conn:ping()
    local err = ffi.lib.mysql.Ping(self.conn)
    if err == nil or #err == 0 then
        return nil
    end
    return err
end

return {
    dial = function(dataSourceName)
        local obj = {
            conn = ffi.lib.mysql.Dial(dataSourceName)
        }
        setmetatable(obj, { __index = mysql_conn })
        return obj
    end
}
