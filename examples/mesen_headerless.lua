load_core("../mesen/Libretro/mesen_libretro.dylib")
load_game("../roms/Nintendo - Nintendo Entertainment System/Yoshi (USA).zip")
for i=1,200 do run() end
local w, h, _, _ = get_video()
print(w, h)
if w ~= 320 or h ~= 240 then error("wrong framebuffer size") end
