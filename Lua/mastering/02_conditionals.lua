
level = 1
num_lives = 5
score = 1000
time_ellapsed = 0

-- TODO: The game starts
if score >= 1000 then
    level = level + 1
    num_lives = 5
else
    time_ellapsed = time_ellapsed + 1
end

print("in level: ".. level)

-- one line conditional
menu_option = 1

if menu_option == 1 then
    menu_text = "Can I play, Daddy?"
elseif menu_option == 2 then
    menu_text = "Don't hurt me"
elseif menu_option == 3 then
    menu_text = "Bring 'em on"
elseif menu_option == 4 then
    menu_text = "I am Death incarnate!"
else
    menu_text = "Rarely"
end

print("difficultuy: ".. menu_text)

