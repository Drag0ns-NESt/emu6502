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

// setZeroPage is used to set result of operation to zeroPage address. It is
// assumed that argument was get from the same address and program counter was
// incremented beforehand
func (cpu *CPU6502) setZeroPage(value uint8) {
	// PC already incremented by 2 (1 for opcode, 1 for zeropage address) so
	// if we want to get the very zeropage address we used initially we should
	// use PC - 1
	cpu.Memory[cpu.Memory[cpu.PC-1]] = value
}

// executeZeroPage is used for operations for performing operations and than
// storing result using zeropage address
func (cpu *CPU6502) executeZeroPage(operation func(value uint8) uint8) {
	cpu.PC += 1

	// Get the address once so that we can use it for both lookup and store
	address := cpu.Memory[cpu.PC]
	cpu.Memory[address] = operation(cpu.Memory[address])
}
