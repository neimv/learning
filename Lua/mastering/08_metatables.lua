
local meta = {}
local vector3d = {}

------------------------------------------------------------------------------
-- declare a new vector3d constructor
------------------------------------------------------------------------------
function vector3d.new(x, y, z)
    local v = {x = x, y = y, z = z}
    setmetatable(v, meta)
    return v
end

function vector3d.add(v1, v2)
    return vector3d.new(v1.x + v2.x, v1.y + v2.y, v1.z + v2.z)
end

function vector3d.substract(v1, v2)
    return vector3d.new(v1.x - v2.x, v1.y - v2.y, v1.z - v2.z)
end

function vector3d.tostring(v)
    return "("..v.x..","..v.y..","..v.z..")"
end
meta.__add = vector3d.add
meta.__sub = vector3d.substract
meta.__tostring = vector3d.tostring

------------------------------------------------------------------------------
-- create two vector3d tables
------------------------------------------------------------------------------
position = vector3d.new(5.0, 5.0, 5.0)
velocity = vector3d.new(10.0, -3.5, 0.0)

local result = velocity + position
local result_subs = velocity - position

print(position)
print(velocity)
print(result)
print(result_subs)

