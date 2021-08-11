load_core(corepath)
load_game(rompath)

print("start")
for i=1,1000 do run() end
local crc = get_fb_crc()
if (crc ~= 3142902422) then error(string.format("wrong page 1, got %d want %d", crc, 3142902422)) end

print("data tests")
set_inputs(0, "1000000000000000")
run()
set_inputs(0, "0000000000000000")
for i=1,400 do run() end
local crc = get_fb_crc()
if (crc ~= 2655120639) then error(string.format("wrong final results, got %d want %d", crc, 2655120639)) end

print("misc tests")
set_inputs(0, "1000000000000000")
run()
set_inputs(0, "0000000000000000")
for i=1,400 do run() end
local crc = get_fb_crc()
if (crc ~= 54675991) then error(string.format("wrong final results, got %d want %d", crc, 54675991)) end

print("sprite test")
set_inputs(0, "1000000000000000")
run()
set_inputs(0, "0000000000000000")
for i=1,400 do run() end
local crc = get_fb_crc()
if (crc ~= 3131102458) then error(string.format("wrong final results, got %d want %d", crc, 3131102458)) end

print("x-scroll latch test")
set_inputs(0, "1000000000000000")
run()
set_inputs(0, "0000000000000000")
for i=1,400 do run() end
local crc = get_fb_crc()
if (crc ~= 1131621047) then error(string.format("wrong final results, got %d want %d", crc, 1131621047)) end

print("hcounter values")
set_inputs(0, "1000000000000000")
run()
set_inputs(0, "0000000000000000")
for i=1,400 do run() end
local crc = get_fb_crc()
if (crc ~= 301684186) then error(string.format("wrong final results, got %d want %d", crc, 301684186)) end

print("end")
set_inputs(0, "1000000000000000")
run()
set_inputs(0, "0000000000000000")
for i=1,400 do run() end
local crc = get_fb_crc()
if (crc ~= 3142902422) then error(string.format("wrong final results, got %d want %d", crc, 3142902422)) end

