package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kivutar/emutest/core"
	"github.com/kivutar/emutest/state"
	"github.com/kivutar/emutest/video"
)

var frames = 0

func runLoop() {
	for frames < state.NFrames {
		// poll inputs here

		state.Core.Run()
		if state.Core.FrameTimeCallback != nil {
			state.Core.FrameTimeCallback.Callback(state.Core.FrameTimeCallback.Reference)
		}
		if state.Core.AudioCallback != nil {
			state.Core.AudioCallback.Callback()
		}

		frames++
	}
}

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.StringVar(&state.CorePath, "L", "", "Path to the libretro core")
	flag.IntVar(&state.NFrames, "nframes", 1, "Number of frames to execute")
	flag.Parse()
	args := flag.Args()

	var gamePath string
	if len(args) > 0 {
		gamePath = args[0]
	}

	vid := video.Init()
	core.Init(vid)

	err := core.Load(state.CorePath)
	if err == nil {
		err := core.LoadGame(gamePath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	runLoop()

	// Unload and deinit in the core.
	core.Unload()
}
