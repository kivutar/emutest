load_core("../mesens/Libretro/mesens_libretro.dylib")
load_game("../roms/Nintendo - Super Nintendo Entertainment System/Super Mario World (USA).zip")
for i=1,200 do run() end
local w, h, _, _ = get_video()
print(w, h)
if w ~= 256 or h ~= 240 then error("wrong framebuffer size") end
