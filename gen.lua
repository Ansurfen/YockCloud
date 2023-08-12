-- mkdir("./dist")
-- for _, value in ipairs({ "etcd", "mongo", "mysql", "postgre", "rabbitmq", "redis" }) do
--     local dst = pathf("./dist", value)
--     mkdir(dst)
--     for _, suffix in ipairs({ "dll", "so", "dylib" }) do
--         alias("suffix", suffix)
--         alias("name", value)
--         sh([[go build -ldflags "-s -w" -o ./dist/$name/$name.$suffix -buildmode=c-shared ./$name/main.go]])
--     end
-- end

---@type golang
local golang = import("opencmd/lang/golang@0.0.1")

golang.mod.tidy()
-- go mod tidy
for _, value in ipairs({ "etcd", "mongo", "mysql", "postgre", "rabbitmq", "redis" }) do
    for _, suffix in ipairs({ "dll", "so", "dylib" }) do
        golang.build({
            buildmode = "c-archive",
            ldflags = {
                s = true,
                w = true
            },
            packages = { strf("./{{.Name}}/main.go", {
                Name = value
            }) },
            output = strf("./dist/{{.Name}}/{{.Name}}.{{.Suffix}}", {
                Suffix = suffix,
                Name = value
            })
        })
    end
end
