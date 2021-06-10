set_options_file("/Users/kivutar/.ludo/mesen_libretro.toml")
load_core("/Users/kivutar/mesen/Libretro/mesen_libretro.dylib")
load_game("/Users/kivutar/roms/Nintendo - Nintendo Entertainment System/Micro Mages.zip")

for i=1,30 do run() end -- skip some frames

load_state("/Users/kivutar/.ludo/savestates/Micro Mages@2021-03-10-13-58-02.state")

if get_sram() ~= "" then error("sram not empty") end

for i=1,10 do
	run()
	local w, h, pitch, fb = get_video()
	print(i, w, h, pitch, string.len(fb))
	if string.len(fb) ~= pitch * h then error("fb not pitch * height") end
end
