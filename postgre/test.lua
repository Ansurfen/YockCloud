---@type comp_postgre
local postgre  = import("./index")

local host     = "192.168.127.128"
local port     = 5432
local user     = "postgres"
local password = "123456"
local dbname   = "postgres"

local conn     = postgre.dial(string.format("host=%s port=%d user=%s " ..
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname))

print(conn:exec("create database bank;"))

-- table.dump(conn:query("select * from bank.customers;"))
