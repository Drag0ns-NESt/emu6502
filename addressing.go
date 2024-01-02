package emu6502

// absoluteAddress returns next operation argument address using absolute
// addressing mode. Updates PC
func (cpu *CPU6502) absoluteAddress() uint16 {
	cpu.PC += 1
	lower := uint16(cpu.Memory[cpu.PC])

	cpu.PC += 1
	higher := uint16(cpu.Memory[cpu.PC])

	return higher<<8 | lower
}

// indirectAddress returns address for next operation using indirect addressing
// mode. Updates PC
func (cpu *CPU6502) indirectAddress() uint16 {
	initialAddress := cpu.absoluteAddress()
	lower := uint16(cpu.Memory[initialAddress])
	higher := uint16(cpu.Memory[initialAddress+1])

	return higher<<8 | lower
}

// indexedIndirectAddress returns address for next operation using indexed
// indirect addressing mode (address, X). Updates PC
func (cpu *CPU6502) indexedIndirectAddress() uint16 {
	// Getting initial address using zeropage, X adrressing
	address := uint16(cpu.zeroPageX())

	// Return indirect address
	return uint16(cpu.Memory[address+1])<<8 + uint16(cpu.Memory[address])
}

// absoluteAddressX returns next operation argument address using absolute, X
// addressing mode, Updates PC
func (cpu *CPU6502) absoluteAddressX() uint16 {
	return cpu.absoluteAddress() + uint16(cpu.X)
}

// absoluteAddressY returns next operation argument address using absolute, Y
// addressing mode, Updates PC
func (cpu *CPU6502) absoluteAddressY() uint16 {
	return cpu.absoluteAddress() + uint16(cpu.Y)
}

// immediate is used to get argument for operation using immediate addressing mode
// just getting value of the byte next to the opcode. Updates PC
func (cpu *CPU6502) immediate() uint8 {
	cpu.PC += 1

	return cpu.Memory[cpu.PC]
}

// absolute is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address. Updates PC
func (cpu *CPU6502) absolute() uint8 {
	return cpu.Memory[cpu.absoluteAddress()]
}

// absoluteX is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address plus X register. Updates PC
func (cpu *CPU6502) absoluteX() uint8 {
	return cpu.Memory[cpu.absoluteAddressX()]
}

// absoluteY is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address plus value in Y register. Updates PC
func (cpu *CPU6502) absoluteY() uint8 {
	return cpu.Memory[cpu.absoluteAddressY()]
}

// indexedIndirect is used to get argument for operation using indexedIndirect ($address, X)
// addressing. Updates PC
func (cpu *CPU6502) indexedIndirect() uint8 {
	return cpu.Memory[cpu.indexedIndirectAddress()]
}

// indirectIndexed is used to get argument for operation using indirect indexed ($address), Y
// addressing. Updates PC
func (cpu *CPU6502) indirectIndexed() uint8 {
	return cpu.Memory[cpu.zeroPage()+cpu.Y]
}

// relative returns relative to current PC address that can be used by controlling
// instructions for jumps. Updates PC
func (cpu *CPU6502) relative() uint16 {
	// 0xff = -1 = 0b1111_1111
	// 0xfe = -2 = 0b1111_1110
	// 0xfd = -3 = 0b1111_1101
	// 0xfc = -4 = 0b1111_1100

	// +2 because we start to count after the instruction ([Instruction] [Relative argument])
	// we assume all controling instructions will take 2 byte
	offset := int16(int8(cpu.Memory[cpu.PC])) + 2

	value := uint16(int16(cpu.PC) + offset)

	cpu.PC += 1

	return value
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

// executeWithAccumulator is used for operations for performing operations and than
// storing result using accumulator register
func (cpu *CPU6502) executeWithAccumulator(operation func(value uint8) uint8) {
	cpu.A = operation(cpu.A)
}

// executeWithAbsolute is used for performing operations and then storing result
// using zeropage address. Updates PC
func (cpu *CPU6502) executeWithAbsolute(operation func(value uint8) uint8) {
	address := cpu.absoluteAddress()
	cpu.Memory[address] = operation(cpu.Memory[address])
}

// executeWithAbsoluteX is used for performing operations and then storing result
// using zeropage address plus X register. Updates PC
func (cpu *CPU6502) executeWithAbsoluteX(operation func(value uint8) uint8) {
	address := cpu.absoluteAddressX()
	cpu.Memory[address] = operation(cpu.Memory[address])
}

// executeWithAbsoluteY is used for performing operations and then storing result
// using zeropage address plus Y register. Updates PC
func (cpu *CPU6502) executeWithAbsoluteY(operation func(value uint8) uint8) {
	address := cpu.absoluteAddressY()
	cpu.Memory[address] = operation(cpu.Memory[address])
}

// executeWithIndexedIndirect is used for performing operations and the stroing result
// using indexed indirect addressing mode (address, X). Update PC
func (cpu *CPU6502) executeWithIndexedIndirect(operation func(uint8) uint8) {
	address := cpu.indexedIndirectAddress()
	cpu.Memory[address] = operation(cpu.Memory[address])
}

// executeWithZeroPage is used for performing operations and than storing result
// using zeropage address. Updates PC
func (cpu *CPU6502) executeWithZeroPage(operation func(value uint8) uint8) {
	cpu.PC += 1

	// Get the address once so that we can use it for both lookup and store
	address := cpu.Memory[cpu.PC]
	cpu.Memory[address] = operation(cpu.Memory[address])
}

// executeWithZeroPageX is used for performing operations and than storing result
// using zeropage, X address. Updates PC
func (cpu *CPU6502) executeWithZeroPageX(operation func(value uint8) uint8) {
	cpu.PC += 1

	address := cpu.Memory[cpu.Memory[cpu.PC]+cpu.X]
	cpu.Memory[address] = operation(cpu.Memory[address])
}
