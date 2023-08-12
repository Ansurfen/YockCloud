-- local docker = import("opencmd/cn/docker")

-- docker.pull("mysql")
-- docker.run({
--     name = "mysql-server",
--     port = "3306:3306",
--     volume = {},
--     env = {
--         MYSQL_ROOT_PASSWORD = "123456"
--     },
--     image = "mysql"
-- })
---@type comp_mysql
local mysql = import("./index")

local conn = mysql.dial("root:123@tcp(192.168.127.128:3306)/sys")
print(conn:exec("use bank;"))
print(conn:ping())
-- print(conn:close())
-- print(conn:exec([[INSERT INTO customers (customerName, ID, telephone, address)
-- VALUES ('test', '123456789012345678', '12345678901', '123 Main St');
-- ]]))

table.dump(conn:query("select * from bank.customers;"))
print(conn:close())
print(conn:ping())
