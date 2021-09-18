-- Use a different seed every time we start
math.randomseed(os.time())

for count = 1, 500 do
    -- Get random enemy position for 800x600 resolution
    enemy_x = math.random(0, 800)
    enemy_y = math.random(0, 600)

    -- Display the two random values
    print("Enemy "..count.." pos: ("..enemy_x..","..enemy_y..")")
end

print("Thank you, goodbye!")

