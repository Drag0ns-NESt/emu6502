package emu6502

// bpl executes BPL instruction on a given address setting program counter to it
// if negative flag is not set.
func (cpu *CPU6502) bpl(address uint16) {
	cpu.PC += 1

	if !cpu.N {
		cpu.PC = address
	}
}

// jsr executes JSR instruction pushing return address to stack and jumping on a
// given address
func (cpu *CPU6502) jsr(address uint16) {
	cpu.PC += 1

	cpu.pushToStack16(cpu.PC)
	cpu.PC = address
}
