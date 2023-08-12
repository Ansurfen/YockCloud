-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@diagnostic disable: undefined-global, undefined-field
ffi.library("mongo", {
    Dial = { "ptr", { "str" } },
    InsertOne = { "str", { "ptr", "str", "str", "str" } },
    InsertMany = { "str", { "ptr", "str", "str", "str" } },
    FindMany = { "str", { "ptr", "str", "str", "str" } },
    FindOne = { "str", { "ptr", "str", "str", "str" } },
    DeleteOne = { "str", { "ptr", "str", "str", "str" } },
    DeleteMany = { "str", { "ptr", "str", "str", "str" } },
    UpdateByID = { "str", { "ptr", "str", "str", "str", "str" } },
    UpdateOne = { "str", { "ptr", "str", "str", "str", "str" } },
    UpdateMany = { "str", { "ptr", "str", "str", "str", "str" } },
})

local mongo_conn = {}

function mongo_conn:insert_one(db, col, data)
    local res = json.decode(ffi.lib.mongo.InsertOne(self.conn, db, col, json.encode(data)))
    return res["data"], res["err"]
end

function mongo_conn:insert_many(db, col, data)
    local res = json.decode(ffi.lib.mongo.InsertMany(self.conn, db, col, json.encode(data)))
    return res["data"], res["err"]
end

function mongo_conn:find_many(db, col, data)
    local res = json.decode(ffi.lib.mongo.FindMany(self.conn, db, col, json.encode(data)))
    return res["data"], res["err"]
end

function mongo_conn:find_one(db, col, data)
    local res = json.decode(ffi.lib.mongo.FindOne(self.conn, db, col, json.encode(data)))
    if #res["data"] > 0 then
        return res["data"][1], res["err"]
    end
    return {}, res["err"]
end

function mongo_conn:delete_one(db, col, data)
    local res = json.decode(ffi.lib.mongo.DeleteOne(self.conn, db, col, json.encode(data)))
    return res["count"], res["err"]
end

function mongo_conn:delete_many(db, col, data)
    local res = json.decode(ffi.lib.mongo.DeleteMany(self.conn, db, col, json.encode(data)))
    return res["count"], res["err"]
end

function mongo_conn:update_by_id(db, col, id, data)
    local res = json.decode(ffi.lib.mongo.UpdateByID(self.conn, db, col, id, json.encode(data)))
    return res["data"], res["err"]
end

function mongo_conn:update_one(db, col, id, data)
    local res = json.decode(ffi.lib.mongo.UpdateOne(self.conn, db, col, id, json.encode(data)))
    return res["data"], res["err"]
end

function mongo_conn:update_many(db, col, id, data)
    local res = json.decode(ffi.lib.mongo.UpdateMany(self.conn, db, col, id, json.encode(data)))
    return res["data"], res["err"]
end

return {
    dial = function(opt)
        local obj = {
            conn = ffi.lib.mongo.Dial(json.encode(opt))
        }
        setmetatable(obj, { __index = mongo_conn })
        return obj
    end
}
