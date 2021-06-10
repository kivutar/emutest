set_options_toml("mesen-s_hle_coprocessor = \"enabled\"")
load_core("../mesens/Libretro/mesens_libretro.dylib")
load_game("../roms/Nintendo - Super Nintendo Entertainment System/Super Mario Kart (Europe).zip")

for i=1,20 do run() end
local _, _, _, frame1 = dump_video()
for i=1,500 do run() end
local _, _, _, frame2 = dump_video()
if frame1 == frame2 then error("hle coproc not working") end
