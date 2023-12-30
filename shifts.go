package emu6502

// asl performs Arithmetic Shift Left operation on the provided value
func (cpu *CPU6502) asl(value uint8) uint8 {
	// coauthored with chatGPT
	// Set the Carry flag if the most significant bit (bit 7) is 1, otherwise clear it
	cpu.C = (value & 0x80) != 0

	// Perform the shift left operation
	shiftedValue := value << 1

	// Update the Zero flag if the result is zero
	cpu.Z = shiftedValue == 0

	// Update the Negative flag based on bit 7 of the shifted value
	cpu.N = (shiftedValue & 0x80) != 0

	// Update program counter
	cpu.PC += 1
	return shiftedValue
}

// rol performs ROL operation rotating byte with a carry flag to left
// (i.e. C=1 |0|1|1|0|1|0|0|0 => C=0 |1|1|0|1|0|0|0|1)
func (cpu *CPU6502) rol(value uint8) uint8 {
	carry := cpu.C
	cpu.C = (value & 0x80) != 0
	value <<= 1
	if carry {
		value += 1
	}

	return value
}
