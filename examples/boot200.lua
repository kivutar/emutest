set_options_string("duckstation_HLE___Enable = \"true\"\nduckstation_CPU___ExecutionMode = \"Interpreter\"")
load_core(corepath)
load_game(rompath)

for i=1,2000 do
	run()
end

print(filename)
screenshot("screenshots/"..filename..".png")
unload_game()
