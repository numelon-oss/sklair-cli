local fs = require("fs")

local tokenise = require("html/tokeniser")
local parse = require("html/parser")

local src = fs.readFileSync("test.html")

local t = tokenise(src)

for _, token in pairs(t) do
    print(token.type, token.name)

    if token.type == "text" then p(token.value) end

    for var, val in pairs(token.props or {}) do
        p(var, val)
    end
end


p()
p()
p()
p()

print(require("json").encode(parse(t)))
