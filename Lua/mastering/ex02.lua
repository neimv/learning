
------- chars
-- Name: Ryu        -> Hadouken
-- Name: Chun Li    -> Lightning kick
-- Name: Guile      -> Sonic boom
-- Name: Honda      -> Hundred hand slap
-- Name: Ken        -> Hadouken
-- Name: Blanka     -> Electric shock

fighter_name = "Ken"
original_name = fighter_name
fighter_name = string.lower(fighter_name)

if fighter_name == "ryu" or fighter_name == "ken" then
    attack_move = "Hadouken"
elseif fighter_name == "chun li" then
    attack_move = "Lightning kick"
elseif fighter_name == "guile" then
    attack_move = "Sonic boom"
elseif fighter_name == "honda" then
    attack_move = "Hundred hand slap"
elseif fighter_name == "blanka" then
    attack_move = "Electric shock"
else
    attack_move = "Error"
end

print(original_name.." attacks with "..attack_move)

