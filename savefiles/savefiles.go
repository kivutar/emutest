// Package savefiles takes care of saving the game SRAM to the filesystem
package savefiles

import (
	"C"
	"errors"
	"io/ioutil"
	"unsafe"

	"github.com/kivutar/emutest/state"
	"github.com/libretro/ludo/libretro"
)

// GetSRAM prints the content of the SRAM
func GetSRAM() []byte {
	len := state.Core.GetMemorySize(libretro.MemorySaveRAM)
	ptr := state.Core.GetMemoryData(libretro.MemorySaveRAM)
	if ptr == nil || len == 0 {
		return []byte{}
	}

	// convert the C array to a go slice
	return C.GoBytes(ptr, C.int(len))
}

// LoadSRAM saves the game SRAM to the filesystem
func LoadSRAM(path string) error {
	len := state.Core.GetMemorySize(libretro.MemorySaveRAM)
	ptr := state.Core.GetMemoryData(libretro.MemorySaveRAM)
	if ptr == nil || len == 0 {
		return errors.New("unable to get SRAM address")
	}

	// this *[1 << 30]byte points to the same memory as ptr, allowing to
	// overwrite this memory
	destination := (*[1 << 30]byte)(unsafe.Pointer(ptr))[:len:len]
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	copy(destination, source)

	return nil
}
