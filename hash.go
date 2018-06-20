package gimli

import (
	"encoding/binary"
)

const rateInBytes = 16

// Hash computes the hash using the sponge construction
func Hash(output, input []byte) {
	var state Gimli
	// absorb full input blocks
	for len(input) >= rateInBytes {
		state[0] ^= binary.LittleEndian.Uint32(input[0:4])
		state[1] ^= binary.LittleEndian.Uint32(input[4:8])
		state[2] ^= binary.LittleEndian.Uint32(input[8:12])
		state[3] ^= binary.LittleEndian.Uint32(input[12:16])
		state.Update()
		input = input[rateInBytes:]
	}

	n := len(input)
	// handle partial input block at the end
	t := make([]byte, 4)
	switch {
	case n > 12:
		state[0] ^= binary.LittleEndian.Uint32(input[0:4])
		state[1] ^= binary.LittleEndian.Uint32(input[4:8])
		state[2] ^= binary.LittleEndian.Uint32(input[8:12])
		copy(t, input[12:])
		state[3] ^= binary.LittleEndian.Uint32(t)
	case n > 8:
		state[0] ^= binary.LittleEndian.Uint32(input[0:4])
		state[1] ^= binary.LittleEndian.Uint32(input[4:8])
		copy(t, input[8:])
		state[2] ^= binary.LittleEndian.Uint32(t)
	case n > 4:
		state[0] ^= binary.LittleEndian.Uint32(input[0:4])
		copy(t, input[4:])
		state[1] ^= binary.LittleEndian.Uint32(t)
	case n > 0:
		copy(t, input)
		state[0] ^= binary.LittleEndian.Uint32(t)
	}

	// do the padding bytes
	paddingIndex := n / 4
	paddingShift := uint((n % 4) * 8)
	state[paddingIndex] ^= 0x1F << (paddingShift)
	// second bit of padding
	state[rateInBytes/4-1] ^= 0x80 << (3 * 8)

	// squeeze full output blocks
	for len(output) >= rateInBytes {
		state.Update()
		binary.LittleEndian.PutUint32(output[0:4], state[0])
		binary.LittleEndian.PutUint32(output[4:8], state[1])
		binary.LittleEndian.PutUint32(output[8:12], state[2])
		binary.LittleEndian.PutUint32(output[12:16], state[3])
		output = output[rateInBytes:]
	}
	n = len(output)
	if n == 0 {
		return
	}
	// handle partial output block
	state.Update()
	switch {
	case n > 12:
		binary.LittleEndian.PutUint32(output[0:4], state[0])
		binary.LittleEndian.PutUint32(output[4:8], state[1])
		binary.LittleEndian.PutUint32(output[8:12], state[2])
		binary.LittleEndian.PutUint32(t, state[3])
		copy(output[12:], t)
	case n > 8:
		binary.LittleEndian.PutUint32(output[0:4], state[0])
		binary.LittleEndian.PutUint32(output[4:8], state[1])
		binary.LittleEndian.PutUint32(t, state[2])
		copy(output[8:], t)
	case n > 4:
		binary.LittleEndian.PutUint32(output[0:4], state[0])
		binary.LittleEndian.PutUint32(t, state[1])
		copy(output[4:], t)
	case n > 0:
		binary.LittleEndian.PutUint32(t, state[0])
		copy(output, t)
	}
}
