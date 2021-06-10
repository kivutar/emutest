// Package video takes care on the game display. It also creates the window
// using GLFW. It exports the Refresh callback used by the libretro
// implementation.
package video

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"unsafe"

	"github.com/libretro/ludo/libretro"
)

var Geom libretro.GameGeometry
var Width, Height, Pitch int32
var PixelFormat uint32
var Rotation uint
var BPP int
var Framebuffer []byte

const (
	// BIT_FORMAT_SHORT_5_5_5_1 has 5 bits R, 5 bits G, 5 bits B, 1 bit alpha
	BIT_FORMAT_SHORT_5_5_5_1 = iota
	// BIT_FORMAT_INT_8_8_8_8_REV has 8 bits R, 8 bits G, 8 bits B, 8 bit alpha
	BIT_FORMAT_INT_8_8_8_8_REV
	// BIT_FORMAT_SHORT_5_6_5 has 5 bits R, 6 bits G, 5 bits
	BIT_FORMAT_SHORT_5_6_5
)

type Format func(data []byte, index int) color.RGBA

// SetPixelFormat is a callback passed to the libretro implementation.
// It allows the core or the game to tell us which pixel format should be used for the display.
func SetPixelFormat(format uint32) bool {
	PixelFormat = format
	switch format {
	case libretro.PixelFormat0RGB1555:
		PixelFormat = format
		BPP = 2
		return true
	case libretro.PixelFormatXRGB8888:
		PixelFormat = format
		BPP = 4
		return true
	case libretro.PixelFormatRGB565:
		PixelFormat = format
		BPP = 2
		return true
	}
	return false
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

func Screenshot(path string) error {
	packedWidth := int(Pitch) / BPP
	img := image.NewRGBA(image.Rect(0, 0, int(Width), int(Height)))
	drawRgbaImage(int(PixelFormat), int(Width), int(Height), packedWidth, BPP, Framebuffer, img)
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	return png.Encode(fd, img)
}

func drawRgbaImage(pixFormat int, w int, h int, packedW int, bpp int, data []byte, dst *image.RGBA) {
	switch pixFormat {
	case BIT_FORMAT_SHORT_5_6_5:
		toRgba(rgb565, w, h, packedW, bpp, data, dst)
	case BIT_FORMAT_INT_8_8_8_8_REV:
		toRgba(rgba8888, w, h, packedW, bpp, data, dst)
	case BIT_FORMAT_SHORT_5_5_5_1:
		fallthrough
	default:
		dst = nil
	}
}

func toRgba(fn Format, w int, h int, packedW int, bpp int, data []byte, dst *image.RGBA) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			index := (y*packedW + x) * bpp
			c := fn(data, index)
			i := (y-dst.Rect.Min.Y)*dst.Stride + (x-dst.Rect.Min.X)*4
			s := dst.Pix[i : i+4 : i+4]
			s[0] = c.R
			s[1] = c.G
			s[2] = c.B
			s[3] = c.A
		}
	}
}

func rgb565(data []byte, index int) color.RGBA {
	pixel := (uint32)(data[index]) + ((uint32)(data[index+1]) << 8)

	return color.RGBA{
		R: byte((pixel >> 8) & 0xF8),
		G: byte((pixel >> 3) & 0xFC),
		B: byte((pixel << 3) & 0xF8),
		A: 255,
	}
}

func rgba8888(data []byte, index int) color.RGBA {
	return color.RGBA{
		R: data[index+2],
		G: data[index+1],
		B: data[index],
		A: 255,
	}
}
