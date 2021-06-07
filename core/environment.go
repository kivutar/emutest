package core

import (
	"fmt"
	"os"
	"os/user"
	"time"
	"unsafe"

	"github.com/kivutar/emutest/options"
	"github.com/kivutar/emutest/state"
	"github.com/libretro/ludo/libretro"
)

var logLevels = map[uint32]string{
	libretro.LogLevelDebug: "DEBUG",
	libretro.LogLevelInfo:  "INFO",
	libretro.LogLevelWarn:  "WARN",
	libretro.LogLevelError: "ERROR",
	libretro.LogLevelDummy: "DUMMY",
}

func logCallback(level uint32, str string) {
	fmt.Printf("[%s]: %s", logLevels[level], str)
}

func getTimeUsec() int64 {
	return time.Now().UnixNano() / 1000
}

func environmentGetVariable(data unsafe.Pointer) bool {
	variable := libretro.GetVariable(data)
	for _, v := range Options.Vars {
		if variable.Key() == v.Key {
			variable.SetValue(v.Choices[v.Choice])
			return true
		}
	}
	return false
}

func environmentSetPixelFormat(data unsafe.Pointer) bool {
	format := libretro.GetPixelFormat(data)
	return vid.SetPixelFormat(format)
}

func environmentGetUsername(data unsafe.Pointer) bool {
	currentUser, err := user.Current()
	if err != nil {
		libretro.SetString(data, "")
	} else {
		libretro.SetString(data, currentUser.Username)
	}
	return true
}

func environmentGetSystemDirectory(data unsafe.Pointer) bool {
	err := os.MkdirAll(state.SystemDirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	libretro.SetString(data, state.SystemDirectory)
	return true
}

func environmentGetSaveDirectory(data unsafe.Pointer) bool {
	err := os.MkdirAll(state.SavefilesDirectory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	libretro.SetString(data, state.SavefilesDirectory)
	return true
}

func environmentSetVariables(data unsafe.Pointer) bool {
	variables := libretro.GetVariables(data)

	pass := []options.VariableInterface{}
	for _, va := range variables {
		va := va
		pass = append(pass, &va)
	}

	var err error
	Options, err = options.New(pass)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func environmentSetCoreOptions(data unsafe.Pointer) bool {
	optionDefinitions := libretro.GetCoreOptionDefinitions(data)

	pass := []options.VariableInterface{}
	for _, cod := range optionDefinitions {
		cod := cod
		pass = append(pass, &cod)
	}

	var err error
	Options, err = options.New(pass)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func environmentSetCoreOptionsIntl(data unsafe.Pointer) bool {
	optionDefinitions := libretro.GetCoreOptionsIntl(data)

	pass := []options.VariableInterface{}
	for _, cod := range optionDefinitions {
		cod := cod
		pass = append(pass, &cod)
	}

	var err error
	Options, err = options.New(pass)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func environment(cmd uint32, data unsafe.Pointer) bool {
	switch cmd {
	case libretro.EnvironmentSetRotation:
		return vid.SetRotation(*(*uint)(data))
	case libretro.EnvironmentGetUsername:
		return environmentGetUsername(data)
	case libretro.EnvironmentGetLogInterface:
		state.Core.BindLogCallback(data, logCallback)
	case libretro.EnvironmentGetPerfInterface:
		state.Core.BindPerfCallback(data, getTimeUsec)
	case libretro.EnvironmentSetFrameTimeCallback:
		state.Core.SetFrameTimeCallback(data)
	case libretro.EnvironmentSetAudioCallback:
		state.Core.SetAudioCallback(data)
	case libretro.EnvironmentGetCanDupe:
		libretro.SetBool(data, true)
	case libretro.EnvironmentSetPixelFormat:
		return environmentSetPixelFormat(data)
	case libretro.EnvironmentGetSystemDirectory:
		return environmentGetSystemDirectory(data)
	case libretro.EnvironmentGetSaveDirectory:
		return environmentGetSaveDirectory(data)
	case libretro.EnvironmentShutdown:
		return true
	case libretro.EnvironmentGetCoreOptionsVersion:
		libretro.SetUint(data, 1)
	case libretro.EnvironmentSetCoreOptions:
		return environmentSetCoreOptions(data)
	case libretro.EnvironmentSetCoreOptionsIntl:
		return environmentSetCoreOptionsIntl(data)
	case libretro.EnvironmentGetVariable:
		return environmentGetVariable(data)
	case libretro.EnvironmentSetVariables:
		return environmentSetVariables(data)
	case libretro.EnvironmentGetVariableUpdate:
		libretro.SetBool(data, Options.Updated)
		Options.Updated = false
	case libretro.EnvironmentSetGeometry:
		vid.Geom = libretro.GetGeometry(data)
	case libretro.EnvironmentSetSystemAVInfo:
		avi := libretro.GetSystemAVInfo(data)
		vid.Geom = avi.Geometry
	case libretro.EnvironmentGetFastforwarding:
		libretro.SetBool(data, false)
	case libretro.EnvironmentGetLanguage:
		libretro.SetUint(data, 0)
	case libretro.EnvironmentGetDiskControlInterfaceVersion:
		libretro.SetUint(data, 0)
	case libretro.EnvironmentSetDiskControlInterface:
		state.Core.SetDiskControlCallback(data)
	default:
		//fmt.Println("[Env]: Not implemented:", cmd)
		return false
	}
	return true
}
