package emu6502

// zeroPage is used to get argument for operation using zero page addressing mode.
// Updates PC
func (cpu *CPU6502) zeroPage() uint8 {
	cpu.PC += 1

	// Getting zero-page address from argument
	return cpu.Memory[cpu.Memory[cpu.PC]]
}

// zeroPageX is used to get argument for operation using (zero page + X) addressing mode.
// Updates PC
func (cpu *CPU6502) zeroPageX() uint8 {
	cpu.PC += 1

	// Getting  zero-page address from argument
	return cpu.Memory[cpu.Memory[cpu.PC]+cpu.X]
}

// indexedIndirect is used to get argument for operation using indexedIndirect
// addressing. Updates PC
func (cpu *CPU6502) indexedIndirect() uint8 {
	// Getting initial address using zeropage, X adrressing
	address := uint16(cpu.zeroPageX())

	// Get the argument address remembering that least significant byte is first
	return cpu.Memory[uint16(cpu.Memory[address+1]<<8)+uint16(cpu.Memory[address])]
}
