// Package savestates takes care of serializing and unserializing the game RAM
// to the host filesystem.
package savestates

import (
	"io/ioutil"
	"os"

	"github.com/kivutar/emutest/state"
)

// Get the current serialized state of the core
func Get() ([]byte, error) {
	s := state.Core.SerializeSize()
	return state.Core.Serialize(s)
}

// Load the state from the filesystem
func Load(path string) error {
	s := state.Core.SerializeSize()
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = state.Core.Unserialize(bytes, s)
	return err
}

// Save the state to the filesystem
func Save(path string) error {
	s := state.Core.SerializeSize()
	bytes, err := state.Core.Serialize(s)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, os.ModePerm)
}
