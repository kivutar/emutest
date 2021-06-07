// Package video takes care on the game display. It also creates the window
// using GLFW. It exports the Refresh callback used by the libretro
// implementation.
package video

import (
	"fmt"
	"unsafe"

	"github.com/libretro/ludo/libretro"

	"crypto/sha1"
)

var Geom libretro.GameGeometry

// SetPixelFormat is a callback passed to the libretro implementation.
// It allows the core or the game to tell us which pixel format should be used for the display.
func SetPixelFormat(format uint32) bool {
	fmt.Printf("[Video]: Set Pixel Format: %v\n", format)
	return true
}

// Refresh the texture framebuffer
func Refresh(data unsafe.Pointer, width int32, height int32, pitch int32) {
	n := height * pitch
	bytes := (*[1 << 30]byte)(data)[:n:n]

	fmt.Printf("[Video]: Refresh: %d %d %d %x\n", width, height, pitch, sha1.Sum(bytes))
}

// SetRotation rotates the game image as requested by the core
func SetRotation(rot uint) bool {
	// limit to valid values (0, 1, 2, 3, which rotates screen by 0, 90, 180 270 degrees counter-clockwise)
	rot = rot % 4

	fmt.Printf("[Video]: Set Rotation: %v\n", rot)

	return true
}
