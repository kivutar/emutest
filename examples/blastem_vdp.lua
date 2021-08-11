load_core("../blastem/blastem_libretro.dylib")
load_game("../VDPFIFOTesting/VDPFIFOTesting.bin")

for i=1,100 do run() end

-- screenshot("screenshots/1.png")

local crc = get_fb_crc()

if (crc ~= 618559741) then error(string.format("wrong page 1, got %d want %d", crc, 618559741)) end

set_inputs(0, "1000000000000000")

run()

set_inputs(0, "0000000000000000")

for i=1,400 do run() end

-- screenshot("screenshots/2.png")

local crc = get_fb_crc()

if (crc ~= 3191250123) then error(string.format("wrong final results, got %d want %d", crc, 3191250123)) end
