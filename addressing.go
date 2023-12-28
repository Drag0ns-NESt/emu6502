package emu6502

// immediate is used to get argument for operation using immediate addressing mode
// just getting value of the byte next to the opcode. Updates PC
func (cpu *CPU6502) immediate() uint8 {
	cpu.PC += 1

	return cpu.Memory[cpu.PC]
}

// absolute is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address. Updates PC
func (cpu *CPU6502) absolute() uint8 {
	cpu.PC += 1
	lower := uint16(cpu.Memory[cpu.PC])

	cpu.PC += 1
	higher := uint16(cpu.Memory[cpu.PC])

	return cpu.Memory[higher<<8|lower]
}

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

// executeWithAccumulator is used for operations for performing operations and than
// storing result using accumulator register
func (cpu *CPU6502) executeWithAccumulator(operation func(value uint8) uint8) {
	cpu.A = operation(cpu.A)
}

// executeWithZeroPage is used for operations for performing operations and than
// storing result using zeropage address
func (cpu *CPU6502) executeWithZeroPage(operation func(value uint8) uint8) {
	cpu.PC += 1

	// Get the address once so that we can use it for both lookup and store
	address := cpu.Memory[cpu.PC]
	cpu.Memory[address] = operation(cpu.Memory[address])
}
