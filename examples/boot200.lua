load_core(corepath)
load_game(rompath)

for i=1,200 do
	run()
end

print(filename)
screenshot("screenshots/"..filename..".png")
unload_game()
