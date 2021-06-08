package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kivutar/emutest/core"
	"github.com/kivutar/emutest/savefiles"
	"github.com/kivutar/emutest/savestates"
	"github.com/kivutar/emutest/state"
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

func runLoop() {
	for state.Frame < state.NFrames {
		// poll inputs here

		run()

		savefiles.DumpSRAM()

		state.Frame++
	}
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.StringVar(&state.CorePath, "L", "", "Path to the libretro core")
	flag.IntVar(&state.SkipFrames, "skip", 0, "Number of frames to skip before any action")
	flag.IntVar(&state.NFrames, "nframes", 1, "Number of frames to execute")
	flag.StringVar(&state.StatePath, "loadstate", "", "Path to a savestate to load right after the skipped frames")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	gamePath := args[0]

	exitOnErr(core.Load(state.CorePath))
	exitOnErr(core.LoadGame(gamePath))

	for i := 0; i < state.SkipFrames; i++ {
		fmt.Print("[Skipping]: ")
		run()
	}

	if state.StatePath != "" {
		exitOnErr(savestates.Load(state.StatePath))
	}

	runLoop()

	core.Unload()
}
