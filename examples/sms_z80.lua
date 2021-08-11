load_core(corepath)
load_game(rompath)

for i=1,50000 do run() end

screenshot("screenshots/z80_1.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3189831392) then error(string.format("wrong crc, got %d want %d", crc, 3189831392)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_2.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 618290548) then error(string.format("wrong crc, got %d want %d", crc, 618290548)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_3.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 742400987) then error(string.format("wrong crc, got %d want %d", crc, 742400987)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_4.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3556014553) then error(string.format("wrong crc, got %d want %d", crc, 3556014553)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_5.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3556014553) then error(string.format("wrong crc, got %d want %d", crc, 3556014553)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_6.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 2819934495) then error(string.format("wrong crc, got %d want %d", crc, 2819934495)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_7.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 360400782) then error(string.format("wrong crc, got %d want %d", crc, 360400782)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_8.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 2819934495) then error(string.format("wrong crc, got %d want %d", crc, 2819934495)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_9.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 360400782) then error(string.format("wrong crc, got %d want %d", crc, 360400782)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_10.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3178521079) then error(string.format("wrong crc, got %d want %d", crc, 3178521079)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_11.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3178521079) then error(string.format("wrong crc, got %d want %d", crc, 3178521079)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_12.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3178521079) then error(string.format("wrong crc, got %d want %d", crc, 3178521079)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_13.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3178521079) then error(string.format("wrong crc, got %d want %d", crc, 3178521079)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_14.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 1830758) then error(string.format("wrong crc, got %d want %d", crc, 1830758)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_15.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 3178521079) then error(string.format("wrong crc, got %d want %d", crc, 3178521079)) end

for i=1,50000 do run() end

screenshot("screenshots/z80_16.png")

local crc = get_fb_crc()
print(crc)

if (crc ~= 2335790675) then error(string.format("wrong crc, got %d want %d", crc, 2335790675)) end

