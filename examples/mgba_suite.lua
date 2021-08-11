load_core("../mgba/mgba_libretro.dylib")
load_game("../suite-v0-r71.gba")

for i=1,100 do run() end

--screenshot("screenshots/1.png")

local crc = get_fb_crc()

if (crc ~= 2427652887) then error(string.format("wrong initial screen, got %d want %d", crc, 2427652887)) end

-- Memory Tests

set_inputs(0, "0000000010000000") -- a
run()
set_inputs(0, "0000000000000000")

for i=1,300 do run() end

--screenshot("screenshots/2.png")

local crc = get_fb_crc()

if (crc ~= 2066056395) then error(string.format("wrong memory tests screen, got %d want %d", crc, 2066056395)) end

-- I/O Tests

reset()
for i=1,100 do run() end

set_inputs(0, "0000010000000000") -- down
run()
set_inputs(0, "0000000010000000") -- a
run()
set_inputs(0, "0000000000000000")

for i=1,300 do run() end

--screenshot("screenshots/3.png")

local crc = get_fb_crc()

if (crc ~= 703508713) then error(string.format("wrong I/O tests screen, got %d want %d", crc, 703508713)) end

-- Timing Tests

reset()
for i=1,100 do run() end

set_inputs(0, "0000010000000000") -- down
run()
set_inputs(0, "0000000000000000")
run()
set_inputs(0, "0000010000000000") -- down
run()
set_inputs(0, "0000000010000000") -- a
run()
set_inputs(0, "0000000000000000")

for i=1,300 do run() end

--screenshot("screenshots/4.png")

local crc = get_fb_crc()

if (crc ~= 764529804) then error(string.format("wrong I/O tests screen, got %d want %d", crc, 764529804)) end

-- Timer count-up Tests

reset()
for i=1,100 do run() end

set_inputs(0, "0000010000000000") -- down
run()
set_inputs(0, "0000000000000000")
run()
set_inputs(0, "0000010000000000") -- down
run()
set_inputs(0, "0000000000000000")
run()
set_inputs(0, "0000010000000000") -- down
run()
set_inputs(0, "0000000010000000") -- a
run()
set_inputs(0, "0000000000000000")

for i=1,300 do run() end

--screenshot("screenshots/5.png")

local crc = get_fb_crc()

if (crc ~= 799453015) then error(string.format("wrong I/O tests screen, got %d want %d", crc, 799453015)) end

