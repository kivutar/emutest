// Package video takes care on the game display. It also creates the window
// using GLFW. It exports the Refresh callback used by the libretro
// implementation.
package video

import (
	"unsafe"

	"github.com/libretro/ludo/libretro"
)

var Geom libretro.GameGeometry
var Width, Height, Pitch int32
var PixelFormat uint32
var Rotation uint
var Framebuffer []byte

// SetPixelFormat is a callback passed to the libretro implementation.
// It allows the core or the game to tell us which pixel format should be used for the display.
func SetPixelFormat(format uint32) bool {
	PixelFormat = format
	return true
}

// Refresh the texture framebuffer
func Refresh(data unsafe.Pointer, width int32, height int32, pitch int32) {
	Width = width
	Height = height
	Pitch = pitch

	n := height * pitch
	Framebuffer = (*[1 << 30]byte)(data)[:n:n]
}

// SetRotation rotates the game image as requested by the core
func SetRotation(rot uint) bool {
	// limit to valid values (0, 1, 2, 3, which rotates screen by 0, 90, 180 270 degrees counter-clockwise)
	Rotation = rot % 4
	return true
}
