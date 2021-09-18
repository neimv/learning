--[[
	Begin a new way
--]]

human = {}
computer = {}

cw = 777
ch = 472

function love.load()
	-- loads once at launch
	love.window.setTitle('DiCE')
	love.window.setMode(cw, ch, {resizable=false, vsync=false})
	love.graphics.setBackgroundColor(0, 0, 0)
	math.randomseed(os.time())
	human.img = love.graphics.newImage('img/digital-die0.png')
	human.win = false
	computer.img = love.graphics.newImage('img/digital-die0.png')
	start = true
	human.name = "Human"
	computer.name = "Computer"
	font = love.graphics.setNewFont("font/orbitron-bold-webfont.ttf", 72)
end

function love.draw()
	-- main loop
	if start == false then
		if human.win == true then
			love.graphics.setColor(20/255, 255/255, 20/255)
			love.graphics.printf("human wins!!", 0, ch-76,cw, 'center')
		else
			love.graphics.setColor(20/255, 255/255, 20/255)
			love.graphics.printf("computer wins!!", 0, ch-76,cw, 'center')
		end
	else
		love.graphics.printf("Click to roll", 0, ch-76,cw, 'center')
	end
	
	love.graphics.draw(human.img, 33, 30, 0, 0.2, 0.2)
	love.graphics.draw(computer.img, cw*0.5, 30, 0, 0.2, 0.2)
end

function love.mousereleased()
	start = false
	computer.roll = math.random(1, 6)
	human.roll = math.random(1, 6)
	
	-- set graphics
	human.img = love.graphics.newImage('img/digital-die'..human.roll..'.png')
	human.img = love.graphics.newImage('img/digital-die'..computer.roll..'.png')
	
	if human.roll > computer.roll then
		human.win = true
	else
		human.win = false
	end
end

