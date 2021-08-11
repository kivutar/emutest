load_core(corepath)
load_game(rompath)

print("skipping bios screen")
for i=1,1000 do
	run()
end

print("detecting digital pad 1")
--screenshot("screenshots/"..filename..".png")
local crc = get_fb_crc()
if (crc ~= 2137335141) then error("could not detect pad 1") end

print("detecting digital pad 2")
set_options_string("duckstation_Controller2___Type = \"DigitalController\"")
for i=1,64 do
	run()
end
--screenshot("screenshots/"..filename.."2.png")
local crc = get_fb_crc()
if (crc ~= 1247281698) then error("could not detect pad 2") end

-- print("detecting analog pad 1")
-- set_options_string("duckstation_Controller1___Type = \"AnalogController\"")
-- for i=1,64 do
-- 	run()
-- end
-- --screenshot("screenshots/"..filename.."3.png")
-- local crc = get_fb_crc()
-- print(crc)

print("detecting pad 1 button presses")
set_inputs(0, "1010100000000000") -- cross, up, select
for i=1,64 do
	run()
end
--screenshot("screenshots/"..filename.."4.png")
local crc = get_fb_crc()
if (crc ~= 4037817268) then error("could not cross, up and select pressed on pad 1") end

print("back to blank")
set_inputs(0, "0000000000000000")
set_inputs(1, "0000000000000000")
for i=1,64 do
	run()
end
--screenshot("screenshots/"..filename.."5.png")
local crc = get_fb_crc()
if (crc ~= 1247281698) then error("could not come back to blank state") end

print("detecting pad 2 button presses")
set_inputs(1, "0101010000000000") -- square, down, start
for i=1,64 do
	run()
end
--screenshot("screenshots/"..filename.."5.png")
local crc = get_fb_crc()
if (crc ~= 1040198715) then error("could not square, down and start pressed on pad 2") end

print("back to blank")
set_inputs(0, "0000000000000000")
set_inputs(1, "0000000000000000")
for i=1,64 do
	run()
end
--screenshot("screenshots/"..filename.."5.png")
local crc = get_fb_crc()
if (crc ~= 1247281698) then error("could not come back to blank state") end

unload_game()
