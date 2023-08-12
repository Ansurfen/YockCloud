-- Copyright 2023 The YockCloud Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

local etcd = import("./index")

---@type etcd_conn
local conn = etcd.dial({
    ["dial-timeout"] = 5 * time.Second,
    endpoints = { "192.168.127.128:2379" }
})

conn:watch({
    key = "testGo",
    callback = function(opt)
        table.dump(opt)
    end
})
