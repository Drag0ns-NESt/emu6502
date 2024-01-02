package emu6502

// sta executes STA (STore Accumulator in memory). Argument is presumed to
// be value stored already in the address, so that it is ignored
func (cpu *CPU6502) sta(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.A
}
