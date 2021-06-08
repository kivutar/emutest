// Package state holds the global state of the app. It is a separate package
// so we can import it from anywhere.
package state

import (
	"os"

	"path/filepath"

	"github.com/libretro/ludo/libretro"
)

// Core is the current libretro core, if any is loaded
var Core *libretro.Core

// Frame is the frame counter
var Frame int

// NFrames is the number of frames to execute
var NFrames int

// SkipFrames is the number of frames to skip before doing stuff
var SkipFrames int

// CorePath is the path of the current libretro core
var CorePath string

// GamePath is the path of the current game
var GamePath string

// StatePath is the path of the savestate to load after frames are skipped
var StatePath string

// OptionsPath is a path to a toml core option file
var OptionsPath string

var HomeDirectory, _ = os.UserHomeDir()

var SavestatesDirectory = filepath.Join(HomeDirectory, "emutest", "savestates")
var SavefilesDirectory = filepath.Join(HomeDirectory, "emutest", "savefiles")
var ScreenshotsDirectory = filepath.Join(HomeDirectory, "emutest", "screenshots")
var SystemDirectory = filepath.Join(HomeDirectory, "emutest", "system")
