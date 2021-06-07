// Package video takes care on the game display. It also creates the window
// using GLFW. It exports the Refresh callback used by the libretro
// implementation.
package video

import (
	"fmt"
	"unsafe"

	"github.com/libretro/ludo/libretro"
)

// Video holds the state of the video package
type Video struct {
	Geom libretro.GameGeometry

	pitch         int32  // pitch set by the refresh callback
	pixFmt        uint32 // format set by the environment callback
	pixType       uint32
	bpp           int32
	width, height int32 // dimensions set by the refresh callback
	rot           uint
}

// Init instanciates the video package
func Init() *Video {
	vid := &Video{}
	vid.Configure()
	return vid
}

// Configure instanciates the video package
func (video *Video) Configure() {
	// Some cores won't call SetPixelFormat, provide default values
	if video.pixFmt == 0 {
		// video.pixFmt = gl.UNSIGNED_SHORT_5_5_5_1
		// video.pixType = gl.BGRA
		// video.bpp = 2
	}
}

// SetPixelFormat is a callback passed to the libretro implementation.
// It allows the core or the game to tell us which pixel format should be used for the display.
func (video *Video) SetPixelFormat(format uint32) bool {
	fmt.Printf("[Video]: Set Pixel Format: %v\n", format)

	// // PixelStorei also needs to be updated whenever bpp changes
	// defer gl.PixelStorei(gl.UNPACK_ROW_LENGTH, video.pitch/video.bpp)

	// switch format {
	// case libretro.PixelFormat0RGB1555:
	// 	video.pixFmt = gl.UNSIGNED_SHORT_5_5_5_1
	// 	video.pixType = gl.BGRA
	// 	video.bpp = 2
	// 	return true
	// case libretro.PixelFormatXRGB8888:
	// 	video.pixFmt = gl.UNSIGNED_INT_8_8_8_8_REV
	// 	video.pixType = gl.BGRA
	// 	video.bpp = 4
	// 	return true
	// case libretro.PixelFormatRGB565:
	// 	video.pixFmt = gl.UNSIGNED_SHORT_5_6_5
	// 	video.pixType = gl.RGB
	// 	video.bpp = 2
	// 	return true
	// default:
	// 	fmt.Printf("Unknown pixel type %v", format)
	// }

	return true
}

// ResetPitch should be called when unloading a game so that the next game won't
// be rendered with the wrong pitch
func (video *Video) ResetPitch() {
	video.pitch = 0
}

// ResetRot should be called when unloading a game so that the next game won't
// be rendered with the wrong rotation
func (video *Video) ResetRot() {
	video.rot = 0
}

// Refresh the texture framebuffer
func (video *Video) Refresh(data unsafe.Pointer, width int32, height int32, pitch int32) {
	video.width = width
	video.height = height
	video.pitch = pitch

	if data == nil {
		return
	}

	fmt.Println("[Video]: Refresh:", width, height, pitch)
}

// SetRotation rotates the game image as requested by the core
func (video *Video) SetRotation(rot uint) bool {
	// limit to valid values (0, 1, 2, 3, which rotates screen by 0, 90, 180 270 degrees counter-clockwise)
	video.rot = rot % 4

	fmt.Printf("[Video]: Set Rotation: %v\n", video.rot)

	return true
}
