load_core(corepath)
load_game(rompath)

for i=1,100 do run() end

local crc1 = get_fb_crc()
--screenshot("screenshots/1.png")

save_state(rompath..".state")

for i=1,1000 do run() end

local crc2 = get_fb_crc()
--screenshot("screenshots/2.png")

load_state(rompath..".state")

local crc3 = get_fb_crc()
--screenshot("screenshots/3.png")

if (crc1 == crc2) then error("same fb1 and fb2") end
if (crc2 ~= crc3) then error("different fb2 and fb3") end

