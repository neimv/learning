-- Use a different seed every time we start
math.randomseed(os.time())

-- Set position of the hero
player_x, player_y = 400, 300
num_enemies = 0

while num_enemies < 500 do
    -- Get random enemy position for 800x600 resolution
    enemy_x = math.random(0, 800)
    enemy_y = math.random(0, 600)

    if (player_x == enemy_x and player_y == enemy_y) then
        print("Enemy and player position clashed")
    else
        -- Display the two random values
        print("Enemy "..num_enemies.. " pos: ("..enemy_x..","..enemy_y..")")
        -- makes sure we increment the number of enemies
        num_enemies = num_enemies + 1
    end
end

print("Thank you, goodbye!")

