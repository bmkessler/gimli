package gimli

import (
	"math/bits"
)

// Gimli stores the state of the gimli permutation
type Gimli [12]uint32

// Update advances the permutation to the next state
func (state *Gimli) Update() {
	for round := 24; round > 0; round-- {
		for column := 0; column < 4; column++ {
			// sp-box
			x := bits.RotateLeft32(state[column], 24)
			y := bits.RotateLeft32(state[4+column], 9)
			z := state[8+column]

			state[8+column] = x ^ (z << 1) ^ ((y & z) << 2)
			state[4+column] = y ^ x ^ ((x | z) << 1)
			state[column] = z ^ y ^ ((x & y) << 3)
		}
		switch round & 3 {
		case 0:
			// small swap
			state[0], state[1], state[2], state[3] = state[1], state[0], state[3], state[2]
			// add constant
			state[0] ^= (0x9e377900 | uint32(round))
		case 2:
			// big swap
			state[0], state[1], state[2], state[3] = state[2], state[3], state[0], state[1]
		}
	}
}
