// Package input exposes the two input callbacks Poll and State needed by the
// libretro implementation. It uses GLFW to access keyboard and joypad, and
// takes care of binding and auto configuring joypads.
package input

import (
	"github.com/libretro/ludo/libretro"
)

// MaxPlayers is the maximum number of players to poll input for
const MaxPlayers = 5

// Hot keys
const (
	// ActionLast is used for iterating
	ActionLast uint32 = libretro.DeviceIDJoypadR3 + 1
)

// States can store the state of inputs for all players
type States [MaxPlayers][ActionLast]int16

// AnalogStates can store the state of analog inputs for all players
type AnalogStates [MaxPlayers][2][2]int16

// Input state for all the players
var (
	NewState       States       // input state for the current frame
	NewAnalogState AnalogStates // analog input state for the current frame
)

// Poll calculates the input state. It is meant to be called for each frame.
func Poll() {
}

func SetState(port uint, values string) {
	for i, char := range values {
		if char == '1' {
			NewState[port][i] = 1
		} else {
			NewState[port][i] = 0
		}
	}
}

// State is a callback passed to core.SetInputState
// It returns 1 if the button corresponding to the parameters is pressed
func State(port uint, device uint32, index uint, id uint) int16 {
	if port >= MaxPlayers {
		return 0
	}

	if device == libretro.DeviceJoypad {
		if id >= uint(ActionLast) || index > 0 {
			return 0 // invalid
		}
		return NewState[port][id]
	}
	if device == libretro.DeviceAnalog {
		if id > 1 || index > 1 {
			return 0 // invalid
		}

		return NewAnalogState[port][index][id]
	}

	return 0
}
