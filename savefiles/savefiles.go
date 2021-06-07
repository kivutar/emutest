// Package savefiles takes care of saving the game SRAM to the filesystem
package savefiles

import (
	"C"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"unsafe"

	"github.com/kivutar/emutest/state"
	"github.com/kivutar/emutest/utils"
	"github.com/libretro/ludo/libretro"
)

var mutex sync.Mutex

// path returns the path of the SRAM file for the current core
func path() string {
	return filepath.Join(
		state.SavefilesDirectory,
		utils.FileName(state.GamePath)+".srm")
}

// SaveSRAM saves the game SRAM to the filesystem
func SaveSRAM() error {
	mutex.Lock()
	defer mutex.Unlock()

	len := state.Core.GetMemorySize(libretro.MemorySaveRAM)
	ptr := state.Core.GetMemoryData(libretro.MemorySaveRAM)
	if ptr == nil || len == 0 {
		return errors.New("unable to get SRAM address")
	}

	// convert the C array to a go slice
	bytes := C.GoBytes(ptr, C.int(len))
	err := os.MkdirAll(state.SavefilesDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	fd, err := os.Create(path())
	if err != nil {
		return err
	}

	_, err = fd.Write(bytes)
	if err != nil {
		fd.Close()
		return err
	}

	err = fd.Close()
	if err != nil {
		return err
	}

	return fd.Sync()
}

// LoadSRAM saves the game SRAM to the filesystem
func LoadSRAM() error {
	mutex.Lock()
	defer mutex.Unlock()

	len := state.Core.GetMemorySize(libretro.MemorySaveRAM)
	ptr := state.Core.GetMemoryData(libretro.MemorySaveRAM)
	if ptr == nil || len == 0 {
		return errors.New("unable to get SRAM address")
	}

	// this *[1 << 30]byte points to the same memory as ptr, allowing to
	// overwrite this memory
	destination := (*[1 << 30]byte)(unsafe.Pointer(ptr))[:len:len]
	source, err := ioutil.ReadFile(path())
	if err != nil {
		return err
	}
	copy(destination, source)

	return nil
}
