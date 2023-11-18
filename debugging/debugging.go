// Package debugging has introspection functions for testing
package debugging

import (
	"C"

	"github.com/kivutar/emutest/state"
	"github.com/libretro/ludo/libretro"
)

// GetRAM prints the content of the SRAM
func GetSystemRAM() []byte {
	len := state.Core.GetMemorySize(libretro.MemorySystemRAM)
	ptr := state.Core.GetMemoryData(libretro.MemorySystemRAM)
	if ptr == nil || len == 0 {
		return []byte{}
	}

	// convert the C array to a go slice
	return C.GoBytes(ptr, C.int(len))
}
