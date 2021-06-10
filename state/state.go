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

// OptionsPath is a path to a toml core option file
var OptionsPath string

// OptionsToml is a toml string containing core options
var OptionsToml string

var HomeDirectory, _ = os.UserHomeDir()
var SavefilesDirectory = filepath.Join(HomeDirectory, "emutest", "savefiles")
var SystemDirectory = filepath.Join(HomeDirectory, "emutest", "system")
