package emu6502

// sta executes STA (STore Accumulator in memory). Argument is presumed to
// be value stored already in the address, so that it is ignored
func (cpu *CPU6502) sta(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.A
}

// stx executes STX (STore X in memory). Argument is presumed to be value
// stored already in the address, so that it is ignored
func (cpu *CPU6502) stx(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.X
}

// sty executes STY (STore Y in memory). Argument is presumed to be value
// stored already in the address, so that it is ignored
func (cpu *CPU6502) sty(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.Y
}

// txa executes TXA (Transfer X to Accumulator) operation
func (cpu *CPU6502) txa() {
	cpu.PC += 1
	cpu.A = cpu.X
}
