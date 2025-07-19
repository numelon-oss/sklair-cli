-- void elements are self closing tags (implied close even if close tag not present)
-- https://developer.mozilla.org/en-US/docs/Glossary/Void_element
local voidElements = {
    area = true,
    base = true,
    br = true,
    col = true,
    embed = true,
    hr = true,
    img = true,
    input = true,
    link = true,
    meta = true,
    param = true,
    source = true,
    track = true,
    wbr = true
}

local function parse(tokens)
    local root = { type = "root", children = {} }
    local stack = { root }

    for _, token in ipairs(tokens) do
        if token.type == "tag_open" then
            local node = {
                type = "element",
                tag = token.name,
                props = token.props,
                children = {}
            }
            table.insert(stack[#stack].children, node)

            -- the fix for void elements was somehow actually easy
            if not voidElements[token.name] then
                table.insert(stack, node)
            end
        elseif token.type == "tag_self" then
            local node = {
                type = "element",
                tag = token.name,
                props = token.props,
                children = {}
            }
            table.insert(stack[#stack].children, node)
        elseif token.type == "text" then
            table.insert(stack[#stack].children, {
                type = "text",
                value = token.value
            })
        elseif token.type == "tag_close" then
            if stack[#stack].tag == token.name then
                table.remove(stack)
            else
                print("mismatched closing tag: " .. token.name)
                -- optional error handling probably later idk
                -- TODO: error handling
                -- probaby just a warning
            end
        end
    end

    return root
end

return parse
