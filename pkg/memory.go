package emu6502

// Memory represents memory with a [0x00000-0x10000] mapping used by 6502 CPU
type Memory interface {
	// Read gets the byte from memory using a given index
	Read(index uint16) uint8
	// Write sets the byte with a given index in memory with a given value
	Write(index uint16, value uint8)
}

type rawArrayMemory struct {
	array [0x10000]uint8
}

// NewRawArrayMemory creates new memory instance that is just 65536 byte array
// under the hood
func NewRawArrayMemory() Memory {
	return &rawArrayMemory{}
}

func (ram *rawArrayMemory) Read(index uint16) uint8 {
	return ram.array[index]
}

func (ram *rawArrayMemory) Write(index uint16, value uint8) {
	ram.array[index] = value
}
