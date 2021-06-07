package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kivutar/emutest/core"
	"github.com/kivutar/emutest/savefiles"
	"github.com/kivutar/emutest/state"
)

func runLoop() {
	for state.Frame < state.NFrames {
		// poll inputs here

		state.Core.Run()
		if state.Core.FrameTimeCallback != nil {
			state.Core.FrameTimeCallback.Callback(state.Core.FrameTimeCallback.Reference)
		}
		if state.Core.AudioCallback != nil {
			state.Core.AudioCallback.Callback()
		}

		savefiles.DumpSRAM()

		state.Frame++
	}
}

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.StringVar(&state.CorePath, "L", "", "Path to the libretro core")
	flag.IntVar(&state.NFrames, "nframes", 1, "Number of frames to execute")
	flag.Parse()
	args := flag.Args()

	gamePath := args[0]

	err := core.Load(state.CorePath)
	if err == nil {
		err := core.LoadGame(gamePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	runLoop()

	core.Unload()
}
