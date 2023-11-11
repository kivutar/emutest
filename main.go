package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"os"
	"slices"

	"github.com/kivutar/emutest/core"
	"github.com/kivutar/emutest/debugging"
	"github.com/kivutar/emutest/input"
	"github.com/kivutar/emutest/options"
	"github.com/kivutar/emutest/savefiles"
	"github.com/kivutar/emutest/savestates"
	"github.com/kivutar/emutest/state"
	"github.com/kivutar/emutest/utils"
	"github.com/kivutar/emutest/video"

	"github.com/Shopify/go-lua"
)

func run() {
	state.Core.Run()
	if state.Core.FrameTimeCallback != nil {
		state.Core.FrameTimeCallback.Callback(state.Core.FrameTimeCallback.Reference)
	}
	if state.Core.AudioCallback != nil {
		state.Core.AudioCallback.Callback()
	}
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func registerFuncs(l *lua.State) {
	l.Register("run", func(l *lua.State) int {
		run()
		return 0
	})
	l.Register("get_ram", func(l *lua.State) int {
		sram := debugging.GetSystemRAM()
		l.PushString(string(sram[:]))
		return 1
	})
	l.Register("get_ram_byte", func(l *lua.State) int {
		offset := lua.CheckInteger(l, 1)
		sram := debugging.GetSystemRAM()
		l.PushInteger(int(sram[offset]))
		return 1
	})
	l.Register("get_sram", func(l *lua.State) int {
		sram := savefiles.GetSRAM()
		l.PushString(string(sram[:]))
		return 1
	})
	l.Register("load_sram", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		err := savefiles.LoadSRAM(path)
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		return 0
	})
	l.Register("get_video", func(l *lua.State) int {
		fb := video.Framebuffer
		l.PushInteger(int(video.Width))
		l.PushInteger(int(video.Height))
		l.PushInteger(int(video.Pitch))
		l.PushString(string(fb[:]))
		return 4
	})
	l.Register("get_fb_crc", func(l *lua.State) int {
		fb := video.Framebuffer
		l.PushUnsigned(uint(crc32.ChecksumIEEE(fb[:])))
		return 1
	})
	l.Register("get_logs", func(l *lua.State) int {
		l.PushString(string(core.Logs[:]))
		return 1
	})
	l.Register("load_state", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		err := savestates.Load(path)
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		return 0
	})
	l.Register("save_state", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		err := savestates.Save(path)
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		return 0
	})
	l.Register("get_state_string", func(l *lua.State) int {
		s, err := savestates.Get()
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		l.PushString(string(s[:]))
		return 1
	})
	l.Register("load_core", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		err := core.Load(path)
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		return 0
	})
	l.Register("load_game", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		err := core.LoadGame(path)
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		return 0
	})
	l.Register("unload_game", func(l *lua.State) int {
		core.UnloadGame()
		return 0
	})
	l.Register("reset", func(l *lua.State) int {
		state.Core.Reset()
		return 0
	})
	l.Register("set_options_file", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		state.OptionsPath = path
		if state.Core != nil {
			core.Options.Load()
			core.Options.Updated = true
		}
		return 0
	})
	l.Register("set_options_string", func(l *lua.State) int {
		toml := lua.CheckString(l, 1)
		state.OptionsToml = toml
		if state.Core != nil {
			core.Options.Load()
			core.Options.Updated = true
		}
		return 0
	})
	l.Register("set_option", func(l *lua.State) int {
		key := lua.CheckString(l, 1)
		value := lua.CheckString(l, 2)
		if state.Core == nil {
			return 0
		}

		keyIndex := slices.IndexFunc(core.Options.Vars, func(v *options.Variable) bool {
			return v.Key == key
		})

		if keyIndex == -1 {
			return 0
		}

		valueIndex := slices.Index(core.Options.Vars[keyIndex].Choices, value)

		if valueIndex != -1 {
			core.Options.Vars[keyIndex].Choice = valueIndex
			core.Options.Updated = true
		}

		return 0
	})
	l.Register("get_option", func(l *lua.State) int {
		key := lua.CheckString(l, 1)
		if state.Core == nil {
			return 0
		}

		keyIndex := slices.IndexFunc(core.Options.Vars, func(v *options.Variable) bool {
			return v.Key == key
		})

		if keyIndex == -1 {
			return 0
		}

		variable := core.Options.Vars[keyIndex]
		l.PushString(variable.Choices[variable.Choice])

		return 1
	})
	l.Register("screenshot", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		err := video.Screenshot(path)
		if err != nil {
			l.PushString(err.Error())
			l.Error()
		}
		return 0
	})
	l.Register("set_inputs", func(l *lua.State) int {
		port := lua.CheckUnsigned(l, 1)
		values := lua.CheckString(l, 2)
		input.SetState(port, values)
		return 0
	})
}

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	corePath := flag.String("L", "", "Path to the libretro core. Optional.")
	romPath := flag.String("r", "", "Path to the ROM. Optional.")
	testPath := flag.String("t", "", "Path to the test lua file")
	testRunner := flag.Bool("T", false, "Test runner mode (script must call os.exit)")
	flag.Parse()
	if !flag.Parsed() {
		fmt.Fprintln(os.Stderr, "Error parsing flags")
		os.Exit(-2)
	}

	l := lua.NewState()
	lua.OpenLibraries(l)
	registerFuncs(l)

	if *corePath != "" {
		if err := lua.DoString(l, "corepath=\""+*corePath+"\""); err != nil {
			exitOnErr(err)
		}
	}

	if *romPath != "" {
		if err := lua.DoString(l, "rompath=\""+*romPath+"\""); err != nil {
			exitOnErr(err)
		}
		if err := lua.DoString(l, "filename=\""+utils.FileName(*romPath)+"\""); err != nil {
			exitOnErr(err)
		}
	}

	if *testPath != "" {
		err := lua.DoFile(l, *testPath)
		if err != nil {
			exitOnErr(err)
		}
	} else {
		fmt.Fprintln(os.Stderr, "No test file specified")
		os.Exit(-2)
	}

	if *testRunner {
		fmt.Fprintln(os.Stderr, "Script did not call os.exit")
		os.Exit(-3)
	}
}
