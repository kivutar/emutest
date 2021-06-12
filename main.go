package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kivutar/emutest/core"
	"github.com/kivutar/emutest/input"
	"github.com/kivutar/emutest/savefiles"
	"github.com/kivutar/emutest/savestates"
	"github.com/kivutar/emutest/state"
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerFuncs(l *lua.State) {
	l.Register("run", func(l *lua.State) int {
		run()
		return 0
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
	l.Register("get_state", func(l *lua.State) int {
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
	l.Register("set_options_file", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		state.OptionsPath = path
		return 0
	})
	l.Register("set_options_string", func(l *lua.State) int {
		toml := lua.CheckString(l, 1)
		state.OptionsToml = toml
		return 0
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
		values := lua.CheckString(l, 1)
		input.SetState(port, values)
		return 0
	})
}

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	l := lua.NewState()
	lua.OpenLibraries(l)
	registerFuncs(l)
	if err := lua.DoFile(l, args[0]); err != nil {
		exitOnErr(err)
	}

	core.Unload()
}
