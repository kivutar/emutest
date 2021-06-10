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
emutest test_coproc.lua
```

With a test file testcoproc.lua:

```
set_options_toml("mesen-s_hle_coprocessor = \"enabled\"")
load_core("../mesens/Libretro/mesens_libretro.dylib")
load_game("../roms/Nintendo - Super Nintendo Entertainment System/Super Mario Kart (Europe).zip")

for i=1,20 do run() end
local _, _, _, frame1 = get_video()
for i=1,500 do run() end
local _, _, _, frame2 = get_video()
if frame1 == frame2 then error("hle coproc not working") end

```
