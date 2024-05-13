--[[
  empezando con la estructura de datos stack
]]

stack = {}

function push(element)
    table.insert(stack, element)
end

function pop()
    if isEmpty() then
        print("stack is empty")
        return
    end

    element = table.remove(stack)

    return element
end

function top()
    if isEmpty() then
        print("stack is empty")
        return
    end

    index = #stack
    element = stack[index]

    return element
end

function isEmpty()
    return #stack == 0
end

function showStack()
    if isEmpty() then
        print("stack is empty")
    end

    for i, element in ipairs(stack) do
        print(i .. " --- " .. element)
    end
end

