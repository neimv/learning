
function display_menu()
    -- Display a menu on the console
    print("+-------------------------------------------")
    print("| Welcome, "..os.date())
    print("+-------------------------------------------")
    print("| 1.Generate random enemy position")
    print("| 2.Distance from enemy to player")
    print("| 3.Get angle from enemy to player")
    print("| 4.Exit")
    print("+-------------------------------------------")
end

function get_distance(enemy_x, enemy_y, player_x, player_y)
    local d = math.sqrt((enemy_x - player_x)^2 + (enemy_y - player_y)^2)
    return d
end

-- log = require "log.log"
math.randomseed(os.time())

-- declaring player position in the middle screen
local player = {
    x = 400,
    y = 300
}
-- local player_x, player_y = 400, 300
-- enemy position
local enemy = {
    x = 0,
    y = 0
}
-- local enemy_x, enemy_y = 0, 0
-- local user_option = 0

-- loop while user option is differente than 4 (exit)
while true do
    display_menu()
    -- read the user option from the keyboard
    print("Please, select your option: ")
    user_option = io.read("*n")

    if user_option == 1 then
        enemy.x = math.random(0, 800)
        enemy.y = math.random(0, 600)

        print("new enemy pos ("..enemy.x..","..enemy.y..")")
    elseif user_option == 2 then
        local d = get_distance(enemy.x, enemy.y, player.x, player.y)
        print("Distance from enemy to player: "..d)
    elseif user_option == 3 then
        local a = math.atan2((enemy.y - player.y), (enemy.x - player.x))
        local a_deg = math.deg(a)
        print("Angle between enemy and player is: "..a_deg)
    elseif user_option == 4 then
        break
    elseif user_option == nil then
        print("insert number")
    else
        print("Error option is not valid")
    end
end

print("Thank you, goodbye!!")

