package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Shopify/go-lua"
	"github.com/kivutar/emutest/core"
	"github.com/kivutar/emutest/savefiles"
	"github.com/kivutar/emutest/savestates"
	"github.com/kivutar/emutest/state"
	"github.com/kivutar/emutest/video"
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
	l.Register("dump_sram", func(l *lua.State) int {
		sram := savefiles.GetSRAM()
		l.PushString(string(sram[:]))
		return 1
	})
	l.Register("dump_video", func(l *lua.State) int {
		fb := video.DumpFramebuffer()
		l.PushInteger(int(video.Width))
		l.PushInteger(int(video.Height))
		l.PushInteger(int(video.Pitch))
		l.PushString(string(fb[:]))
		return 4
	})
	l.Register("load_state", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		savestates.Load(path)
		return 0
	})
	l.Register("load_core", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		core.Load(path)
		return 0
	})
	l.Register("load_game", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		core.LoadGame(path)
		return 0
	})
	l.Register("set_options_file", func(l *lua.State) int {
		path := lua.CheckString(l, 1)
		state.OptionsPath = path
		return 0
	})
	l.Register("set_options_toml", func(l *lua.State) int {
		toml := lua.CheckString(l, 1)
		state.OptionsToml = toml
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
