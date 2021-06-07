// Package core takes care of instanciating the libretro core, setting the
// input, audio, video, environment callbacks needed to play the games.
// It also deals with core options and persisting SRAM periodically.
package core

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kivutar/emutest/audio"
	"github.com/kivutar/emutest/input"
	"github.com/kivutar/emutest/options"
	"github.com/kivutar/emutest/state"
	"github.com/kivutar/emutest/video"
	"github.com/libretro/ludo/libretro"
)

// Options holds the settings for the current core
var Options *options.Options

// Load loads a libretro core
func Load(sofile string) error {
	// In case the a core is already loaded, we need to close it properly
	// before loading the new core
	Unload()

	// This must be set before the environment callback is called
	state.CorePath = sofile

	var err error
	state.Core, err = libretro.Load(sofile)
	if err != nil {
		return err
	}
	state.Core.SetEnvironment(environment)
	state.Core.Init()
	state.Core.SetVideoRefresh(video.Refresh)
	state.Core.SetInputPoll(input.Poll)
	state.Core.SetInputState(input.State)
	state.Core.SetAudioSample(audio.Sample)
	state.Core.SetAudioSampleBatch(audio.SampleBatch)

	si := state.Core.GetSystemInfo()
	if len(si.LibraryName) > 0 {
		fmt.Println("[Core]: Name:", si.LibraryName)
		fmt.Println("[Core]: Version:", si.LibraryVersion)
		fmt.Println("[Core]: Valid extensions:", si.ValidExtensions)
		fmt.Println("[Core]: Need fullpath:", si.NeedFullpath)
		fmt.Println("[Core]: Block extract:", si.BlockExtract)
	}

	return nil
}

// unzipGame unzips a rom to tmpdir and returns the path and size of the extracted ROM.
// In case the zip contains more than one file, they are all extracted and the
// first one is passed to the libretro core.
func unzipGame(filename string) (string, int64, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return "", 0, err
	}
	defer r.Close()

	var mainPath string
	var mainSize int64
	for i, cf := range r.File {
		size := int64(cf.UncompressedSize)
		rc, err := cf.Open()
		if err != nil {
			return "", 0, err
		}

		path := filepath.Join(os.TempDir(), cf.Name)

		if cf.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return "", 0, err
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, cf.Mode())
		if err != nil {
			return "", 0, err
		}

		if _, err = io.Copy(outFile, rc); err != nil {
			return "", 0, err
		}
		outFile.Close()
		rc.Close()

		if i == 0 {
			mainPath = path
			mainSize = size
		}
	}

	return mainPath, mainSize, nil
}

// LoadGame loads a game. A core has to be loaded first.
func LoadGame(gamePath string) error {
	if _, err := os.Stat(gamePath); os.IsNotExist(err) {
		return err
	}

	// If we're loading a new game on the same core, save the RAM of the previous
	// game before closing it.
	if state.GamePath != gamePath {
		UnloadGame()
	}

	si := state.Core.GetSystemInfo()

	gi, err := getGameInfo(gamePath, si.BlockExtract)
	if err != nil {
		return err
	}

	if !si.NeedFullpath {
		bytes, err := ioutil.ReadFile(gi.Path)
		if err != nil {
			return err
		}

		gi.SetData(bytes)
	}

	ok := state.Core.LoadGame(*gi)
	if !ok {
		return errors.New("failed to load the game")
	}

	avi := state.Core.GetSystemAVInfo()

	video.Geom = avi.Geometry

	if state.Core.AudioCallback != nil {
		state.Core.AudioCallback.SetState(true)
	}

	state.GamePath = gamePath

	state.Core.SetControllerPortDevice(0, libretro.DeviceJoypad)
	state.Core.SetControllerPortDevice(1, libretro.DeviceJoypad)
	state.Core.SetControllerPortDevice(2, libretro.DeviceJoypad)
	state.Core.SetControllerPortDevice(3, libretro.DeviceJoypad)
	state.Core.SetControllerPortDevice(4, libretro.DeviceJoypad)

	fmt.Println("[Core]: Game loaded: " + gamePath)
	//savefiles.LoadSRAM()

	return nil
}

// Unload unloads a libretro core
func Unload() {
	if state.Core != nil {
		UnloadGame()
		state.Core.Deinit()
		state.CorePath = ""
		state.Core = nil
		Options = nil
	}
}

// UnloadGame unloads a game.
func UnloadGame() {
	//savefiles.SaveSRAM()
	state.Core.UnloadGame()
	state.GamePath = ""
}

// getGameInfo opens a rom and return the libretro.GameInfo needed to launch it
func getGameInfo(filename string, blockExtract bool) (*libretro.GameInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if filepath.Ext(filename) == ".zip" && !blockExtract {
		path, size, err := unzipGame(filename)
		if err != nil {
			return nil, err
		}
		return &libretro.GameInfo{Path: path, Size: size}, nil
	}
	return &libretro.GameInfo{Path: filename, Size: fi.Size()}, nil
}
