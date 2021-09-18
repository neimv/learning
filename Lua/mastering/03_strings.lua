color = "#ce10e3"

-- substiture all occurences of # with nothing
pure_color = string.upper(string.gsub(color, "#", ""))

-- get a substring of a string
pure_color_2 = string.sub(color, 2, string.len(color))

print("The color is " .. pure_color.. " " .. pure_color_2 .. " size of color: " .. #color)

-- multiline string
sea_level = [[
prueba
de las cosas
mas raras
de esto
]]

print(sea_level)

-- Multiple assignment
-- x = 0
-- y = 0
x, y = 0, 0
score, num_live, time_ellapsed = 0, 5, 0

