// Package savestates takes care of serializing and unserializing the game RAM
// to the host filesystem.
package savestates

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kivutar/emutest/state"
)

// Save the current state to the filesystem. name is the name of the
// savestate file to save to, without extension.
func Save(name string) error {
	s := state.Core.SerializeSize()
	bytes, err := state.Core.Serialize(s)
	if err != nil {
		return err
	}
	path := filepath.Join(state.SavestatesDirectory, name+".state")
	err = os.MkdirAll(state.SavestatesDirectory, os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0644)
}

// Load the state from the filesystem
func Load(path string) error {
	fmt.Println("[Savestates]: Loading", path)
	s := state.Core.SerializeSize()
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = state.Core.Unserialize(bytes, s)
	return err
}
