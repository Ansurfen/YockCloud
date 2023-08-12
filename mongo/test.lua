local mongo = import("./index")
local conn = mongo.dial({
    url = "mongodb://192.168.127.128:27017",
    cred = {
        username = "root",
        password = "1234"
    }
})

-- conn:insert_one("testing", "users", {
--     fullName = "User 1",
--     age = 30
-- })

local ids = conn:insert_many("testing", "users", {
    {
        fullName = "User 1",
        age = 30
    },
    {
        fullName = "User 2",
        age = 20
    }
})
table.dump(ids)
for _, id in ipairs(ids) do
    local res = conn:update_by_id("testing", "users", id, {
        ["$set"] = {
            fullName = "User V"
        },
        ["$inc"] = {
            age = 1
        }
    })
    table.dump(res)
end

local d = conn:find_many("testing", "users", {
    ["$and"] = {
        {
            age = {
                ["$gt"] = 10
            }
        }
    }
})
-- table.dump(d)

local cnt = conn:delete_many("testing", "users", {
    ["$and"] = {
        {
            age = {
                ["$gt"] = 0
            }
        }
    }
})
print(cnt)
