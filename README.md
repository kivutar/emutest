# emutest

A simple test framework for libretro cores.

`emutest` can run a certain number of frames, and dump data like the content of the video framebuffer, or the system RAM.

## Setup

```
go install github.com/kivutar/emutest@latest
```

## Usage:

Example:

```
emutest -skip 10 -nframes 5 -loadstate game.state -L fceumm_libretro.so game.nes
```

Usage:

```
Usage of ./emutest:
  -L string
    	Path to the libretro core
  -loadstate string
    	Path to a savestate to load right after the skipped frames
  -nframes int
    	Number of frames to execute (default 1)
  -skip int
    	Number of frames to skip before any action
```
