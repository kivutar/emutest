set_options("/Users/kivutar/.ludo/mesens_libretro.toml")
load_core("/Users/kivutar/mesens/Libretro/mesens_libretro.dylib")
load_game("/Users/kivutar/roms/Nintendo - Super Nintendo Entertainment System/Super Mario Kart (Europe).zip")

for i=1,20 do run() end
local _, _, _, frame1 = dump_video()
for i=1,500 do run() end
local _, _, _, frame2 = dump_video()
if frame1 == frame2 then error("hle coproc not working") end
